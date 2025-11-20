import { ArrowUp, ArrowDown } from 'lucide-react'

interface Transaction {
    id: string
    description: string
    category: string
    amount: number
    type: 'income' | 'expense'
    date: string
}

interface RecentTransactionsProps {
    transactions: Transaction[]
}

export function RecentTransactions({ transactions }: RecentTransactionsProps) {
    return (
        <div className="bg-white rounded-xl border border-border shadow-sm overflow-hidden">
            <div className="p-6 border-b border-border">
                <h3 className="text-lg font-semibold text-foreground">最近の取引</h3>
            </div>
            <div className="divide-y divide-border">
                {transactions.map((tx) => (
                    <div key={tx.id} className="p-4 flex items-center justify-between hover:bg-muted/50 transition-colors">
                        <div className="flex items-center gap-3">
                            <div className={`p-2 rounded-lg ${tx.type === 'income' ? 'bg-accent/10' : 'bg-destructive/10'
                                }`}>
                                {tx.type === 'income' ? (
                                    <ArrowUp className="w-4 h-4 text-accent" />
                                ) : (
                                    <ArrowDown className="w-4 h-4 text-destructive" />
                                )}
                            </div>
                            <div>
                                <p className="font-medium text-foreground">{tx.description}</p>
                                <p className="text-xs text-muted-foreground">{tx.category}</p>
                            </div>
                        </div>
                        <div className="text-right">
                            <p className={`font-semibold ${tx.type === 'income' ? 'text-accent' : 'text-destructive'
                                }`}>
                                {tx.type === 'income' ? '+' : '-'}{tx.amount.toLocaleString()}
                            </p>
                            <p className="text-xs text-muted-foreground">{tx.date}</p>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    )
}
