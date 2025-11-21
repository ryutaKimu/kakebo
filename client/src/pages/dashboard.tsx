import { AppHeader } from '@/components/appHeader'
import { AppSidebar } from '@/components/appSideBar'
import { SummaryCard } from '@/components/summaryCard'
import { TrendingUp, TrendingDown, Plus, Zap, PiggyBank } from 'lucide-react'
import { Link, useNavigate } from 'react-router-dom'
import { useEffect, useState } from 'react'
import { fetchUserFinancialData } from '@/api/kakebo'
import axios from 'axios'
import { API_ERROR } from '@/frontUtils/constants'
import { SavingsGoalCard } from '@/components/savingGoalCard'
export default function DashboardPage() {
    const navigate = useNavigate()
    const [totalIncome, setTotalIncome] = useState<number>(0)
    const [totalExpense, setTotalExpense] = useState<number>(0)
    const [totalSaving, setTotalSaving] = useState<number>(0)
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState<string | null>(null)

    const goals = [
        {
            id: "1",
            name: "ノートパソコン",
            targetAmount: 150000,
            currentAmount: 95000,
            targetDate: "2025-12-31",
            image: "/m1pro_00.jpg",
            purchased: false,
        },
        {
            id: "2",
            name: "海外旅行",
            targetAmount: 500000,
            currentAmount: 500000,
            targetDate: "2026-06-30",
            image: "/holidayTop01_mv.jpg",
            purchased: true,
        },
        {
            id: "3",
            name: "カメラ",
            targetAmount: 200000,
            currentAmount: 120000,
            targetDate: "2026-02-28",
            image: "/camera-photography.png",
            purchased: false,
        },
    ]

    const displayedGoals = goals.slice(0, 2)

    const handleDeleteGoal = (id: string) => {
        console.log("[v0] 目標を削除:", id)
        // 実装: 目標削除処理
    }

    useEffect(() => {
        const fetchData = async () => {
            setIsLoading(true);
            setError(null);
            try {
                const data = await fetchUserFinancialData()
                setTotalIncome(data.total_income)
                setTotalExpense(data.total_cost)
                setTotalSaving(data.saving_amount)
            } catch (err) {
                if (axios.isAxiosError(err) && err.response?.status === 401) {
                    navigate("/login")
                    return
                }
                setError(API_ERROR.FETCH_DATA)
            } finally {
                setIsLoading(false)
            }
        }
        fetchData()
    }, [navigate])

    if (isLoading) {
        // ローディング中のUIを返す
        return <div>読み込み中...</div>;
    }


    if (error) {
        // エラーUIを返す
        return <div>{error}</div>;
    }
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
                            <SummaryCard
                                title="貯蓄目標"
                                amount={`¥${totalSaving.toLocaleString()}`}
                                amountColor="primary"
                                icon={<PiggyBank className="w-5 h-5" />}
                            />
                        </div>


                        <div>
                            <div className="flex items-center justify-between mb-4">
                                <h2 className="text-xl font-semibold text-foreground">欲しいもの</h2>
                                <Link to="/wishlist" className="text-sm text-primary hover:underline">
                                    すべて見る →
                                </Link>
                            </div>
                            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                                {displayedGoals.map((goal) => (
                                    <SavingsGoalCard key={goal.id} goal={goal} onDelete={handleDeleteGoal} />
                                ))}
                            </div>
                        </div>

                        {/* クイックアクション */}
                        <footer className="border-t border-border bg-secondary p-6">
                            <div className="max-w-6xl mx-auto grid grid-cols-2 md:grid-cols-4 gap-3">
                                <Link
                                    to="/income/new"
                                    className="flex items-center justify-center gap-2 bg-primary text-primary-foreground rounded-lg py-3 px-4 font-semibold hover:opacity-90 transition-opacity"
                                >
                                    <TrendingUp className="w-5 h-5" />
                                    <span>収入を記録</span>
                                </Link>
                                <Link
                                    to="/expense/new"
                                    className="flex items-center justify-center gap-2 bg-destructive text-destructive-foreground rounded-lg py-3 px-4 font-semibold hover:opacity-90 transition-opacity"
                                >
                                    <TrendingDown className="w-5 h-5" />
                                    <span>支出を記録</span>
                                </Link>
                                <Link
                                    to="/adjustment/new"
                                    className="flex items-center justify-center gap-2 bg-accent text-accent-foreground rounded-lg py-3 px-4 font-semibold hover:opacity-90 transition-opacity"
                                >
                                    <Zap className="w-5 h-5" />
                                    <span>収益調整</span>
                                </Link>
                                <Link
                                    to="/wishlist/new"
                                    className="flex items-center justify-center gap-2 bg-muted text-foreground rounded-lg py-3 px-4 font-semibold hover:opacity-90 transition-opacity"
                                >
                                    <Plus className="w-5 h-5" />
                                    <span>欲しいもの登録</span>
                                </Link>
                            </div>
                        </footer>
                    </div>
                </main>
            </div>
        </div>
    )
}
