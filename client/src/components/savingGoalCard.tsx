"use client"

import { PieChart, Pie, Cell, ResponsiveContainer, Legend, Tooltip } from "recharts"
import { Calendar, Trash2, Check } from "lucide-react"

interface SavingsGoal {
    id: string
    name: string
    targetAmount: number
    currentAmount: number
    targetDate: string
    image?: string
    purchased?: boolean
}

interface SavingsGoalCardProps {
    goal: SavingsGoal
    onDelete?: (id: string) => void
}

export function SavingsGoalCard({ goal, onDelete }: SavingsGoalCardProps) {
    const displayCurrent = goal.purchased ? goal.targetAmount : goal.currentAmount
    const isExceeded = displayCurrent >= goal.targetAmount
    const remaining = isExceeded ? 0 : goal.targetAmount - displayCurrent
    const percentage = Math.round((displayCurrent / goal.targetAmount) * 100)
    { console.log(goal.image) }

    const chartData = isExceeded
        ? [{ name: "貯金済み", value: goal.targetAmount, fill: "#22C55E" }]
        : [
            { name: "貯金済み", value: displayCurrent, fill: "#3B82F6" },
            { name: "残り", value: remaining, fill: "#E5E7EB" },
        ]

    const formattedDate = new Date(goal.targetDate).toLocaleDateString("ja-JP", {
        year: "numeric",
        month: "2-digit",
        day: "2-digit",
    })

    return (
        <div
            className={`rounded-xl border shadow-sm overflow-hidden hover:shadow-md transition-shadow ${goal.purchased ? "bg-green-50 border-green-200" : "bg-white border-border"
                }`}
        >
            {goal.purchased && (
                <div className="absolute top-3 right-3 bg-green-600 text-white rounded-full p-1.5 z-10">
                    <Check className="w-4 h-4" />
                </div>
            )}

            {/* 画像セクション */}
            {goal.image && (
                <div className={`w-full h-40 bg-muted overflow-hidden relative ${goal.purchased ? "opacity-75" : ""}`}>
                    <img src={goal.image || "/placeholder.svg"} alt={goal.name} className="w-full h-full object-cover" />
                </div>
            )}

            <div className="p-6 space-y-4">
                {/* タイトルと削除ボタン */}
                <div className="flex items-start justify-between">
                    <div>
                        <h3 className="font-semibold text-lg text-foreground">{goal.name}</h3>
                        <p className={`text-sm mt-1 ${goal.purchased ? "text-green-600" : "text-muted-foreground"}`}>
                            {goal.purchased ? "購入済み" : `目標金額: ¥${goal.targetAmount.toLocaleString()}`}
                        </p>
                    </div>
                    {onDelete && (
                        <button
                            onClick={() => onDelete(goal.id)}
                            className="p-2 hover:bg-muted rounded-lg transition-colors"
                            aria-label="削除"
                        >
                            <Trash2 className="w-4 h-4 text-muted-foreground" />
                        </button>
                    )}
                </div>

                {/* 日付情報 */}
                <div className="flex items-center gap-2 text-sm text-muted-foreground">
                    <Calendar className="w-4 h-4" />
                    <span>目標日: {formattedDate}</span>
                </div>

                <div className="grid grid-cols-3 gap-3 py-3 border-y border-border">
                    <div className="text-center">
                        <p className={`text-2xl font-bold ${isExceeded ? "text-green-600" : "text-primary"}`}>
                            {isExceeded ? "達成" : `${percentage}%`}
                        </p>
                        <p className="text-xs text-muted-foreground">達成度</p>
                    </div>
                    <div className="text-center">
                        <p className="text-lg font-semibold text-foreground">¥{displayCurrent.toLocaleString()}</p>
                        <p className="text-xs text-muted-foreground">貯金済み</p>
                    </div>
                    <div className="text-center">
                        <p className={`text-lg font-semibold ${isExceeded ? "text-green-600" : "text-red-600"}`}>
                            ¥{remaining.toLocaleString()}
                        </p>
                        <p className="text-xs text-muted-foreground">{isExceeded ? "達成" : "残り"}</p>
                    </div>
                </div>

                {/* 円グラフ */}
                <div className="h-48 flex items-center justify-center">
                    <ResponsiveContainer width="100%" height="100%">
                        <PieChart>
                            <Pie
                                data={chartData}
                                cx="50%"
                                cy="50%"
                                innerRadius={50}
                                outerRadius={80}
                                paddingAngle={2}
                                dataKey="value"
                            >
                                {chartData.map((entry, index) => (
                                    <Cell key={`cell-${index}`} fill={entry.fill} />
                                ))}
                            </Pie>
                            <Tooltip
                                formatter={(value: number) => `¥${value.toLocaleString()}`}
                                contentStyle={{
                                    backgroundColor: "#fff",
                                    border: "1px solid #e5e7eb",
                                    borderRadius: "0.5rem",
                                }}
                            />
                            {!isExceeded && (
                                <Legend
                                    verticalAlign="bottom"
                                    height={36}
                                    formatter={(value) => {
                                        if (value === "貯金済み") return "✓ 貯金済み"
                                        return "◯ 残り"
                                    }}
                                />
                            )}
                        </PieChart>
                    </ResponsiveContainer>
                </div>
            </div>
        </div>
    )
}
