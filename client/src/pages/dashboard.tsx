import { AppHeader } from '@/components/appHeader'
import { AppSidebar } from '@/components/appSideBar'
import { SummaryCard } from '@/components/summaryCard'
import { ProgressBar } from '@/components/progressBar'
import { TrendingUp, TrendingDown, PiggyBank } from 'lucide-react'
import { RecentTransactions } from '@/components/recentTransactions'
import { Link } from 'react-router-dom'
import { useEffect, useState } from 'react'
import { fetchUserFinancialData } from '@/api/kakebo'
export default function DashboardPage() {
    const [totalIncome, setTotalIncome] = useState<number>(0)
    const [totalExpense, setTotalExpense] = useState<number>(0)
    const savingsGoal = { current: 245000, target: 500000 }
    const savingsPercentage = Math.round((savingsGoal.current / savingsGoal.target) * 100)

    useEffect(() => {
        const fetchData = async () => {
            try {
                const token = localStorage.getItem("access_token")
                if (!token) throw new Error("トークンがありません")
                const data = await fetchUserFinancialData(token)
                setTotalIncome(data.total_income)
                setTotalExpense(data.total_cost)
                console.log(data)
            } catch (err) {
                console.error("データ取得失敗:", err)
            }
        }
        fetchData()
    }, [])
    const recentTransactions = [
        { id: '1', description: '給料', category: '給与', amount: 150000, type: 'income' as const, date: '2025-11-15' },
        { id: '2', description: 'スーパー', category: '食費', amount: 5200, type: 'expense' as const, date: '2025-11-14' },
        { id: '3', description: 'カフェ', category: '外食', amount: 1500, type: 'expense' as const, date: '2025-11-13' },
        { id: '4', description: 'ボーナス', category: '給与', amount: 50000, type: 'income' as const, date: '2025-11-10' },
        { id: '5', description: '家賃', category: '住まい', amount: 80000, type: 'expense' as const, date: '2025-11-01' },
    ]

    return (
        <div className="flex h-screen bg-background">
            <AppSidebar />
            <div className="flex-1 flex flex-col overflow-hidden">
                <AppHeader />
                <main className="flex-1 overflow-y-auto">
                    <div className="max-w-6xl mx-auto p-6 space-y-8">
                        {/* ページタイトル */}
                        <div>
                            <h1 className="text-3xl font-bold text-foreground">ダッシュボード</h1>
                            <p className="text-muted-foreground mt-1">あなたの家計情報をひと目で確認</p>
                        </div>

                        {/* 概要カード */}
                        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                            <SummaryCard
                                title="今月の収入"
                                amount={`¥${totalIncome.toLocaleString()}`}
                                amountColor="accent"
                                icon={<TrendingUp className="w-5 h-5" />}
                            />
                            <SummaryCard
                                title="今月の支出"
                                amount={`¥${totalExpense.toLocaleString()}`}
                                amountColor="destructive"
                                icon={<TrendingDown className="w-5 h-5" />}
                            />
                        </div>

                        {/* 貯金目標 */}
                        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
                            <div className="lg:col-span-2">
                                <ProgressBar
                                    title="貯金目標"
                                    current={savingsGoal.current}
                                    target={savingsGoal.target}
                                    percentage={savingsPercentage}
                                />
                            </div>
                            <Link to="/goals" className="bg-secondary rounded-xl p-6 border border-border shadow-sm flex items-center justify-center hover:shadow-md transition-shadow">
                                <div className="text-center">
                                    <p className="text-muted-foreground mb-2">目標を管理する</p>
                                    <p className="font-semibold text-primary">設定を見る →</p>
                                </div>
                            </Link>
                        </div>

                        {/* 最近の取引 */}
                        <RecentTransactions transactions={recentTransactions} />

                        {/* クイックアクション */}
                        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                            <Link to="/income" className="bg-white rounded-lg p-4 border border-border hover:shadow-md transition-shadow text-center">
                                <div className="inline-flex p-3 bg-accent/10 rounded-lg mb-2">
                                    <TrendingUp className="w-5 h-5 text-accent" />
                                </div>
                                <p className="font-semibold text-foreground">収入を記録</p>
                            </Link>
                            <Link to="/expense" className="bg-white rounded-lg p-4 border border-border hover:shadow-md transition-shadow text-center">
                                <div className="inline-flex p-3 bg-destructive/10 rounded-lg mb-2">
                                    <TrendingDown className="w-5 h-5 text-destructive" />
                                </div>
                                <p className="font-semibold text-foreground">支出を記録</p>
                            </Link>
                            <Link to="/goals" className="bg-white rounded-lg p-4 border border-border hover:shadow-md transition-shadow text-center">
                                <div className="inline-flex p-3 bg-primary/10 rounded-lg mb-2">
                                    <PiggyBank className="w-5 h-5 text-primary" />
                                </div>
                                <p className="font-semibold text-foreground">目標管理</p>
                            </Link>
                            <Link to="/reports" className="bg-white rounded-lg p-4 border border-border hover:shadow-md transition-shadow text-center">
                                <div className="inline-flex p-3 bg-muted rounded-lg mb-2">
                                    <svg className="w-5 h-5 text-muted-foreground" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                                    </svg>
                                </div>
                                <p className="font-semibold text-foreground">レポート</p>
                            </Link>
                        </div>
                    </div>
                </main>
            </div>
        </div>
    )
}
