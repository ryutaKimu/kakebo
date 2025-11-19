import React from 'react'

interface SummaryCardProps {
    title: string
    amount: string
    amountColor?: 'primary' | 'accent' | 'destructive'
    icon?: React.ReactNode
}

export function SummaryCard({ title, amount, amountColor = 'primary', icon }: SummaryCardProps) {
    const colorClasses = {
        primary: 'text-primary',
        accent: 'text-accent',
        destructive: 'text-destructive',
    }

    return (
        <div className="bg-white rounded-xl p-6 border border-border shadow-sm">
            <div className="flex items-center justify-between mb-4">
                <h3 className="text-sm font-medium text-muted-foreground">{title}</h3>
                {icon && <div className="text-muted-foreground">{icon}</div>}
            </div>
            <p className={`text-2xl font-bold ${colorClasses[amountColor]}`}>
                {amount}
            </p>
        </div>
    )
}
