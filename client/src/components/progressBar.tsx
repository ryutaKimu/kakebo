interface ProgressBarProps {
    title: string
    current: number
    target: number
    percentage: number
}

export function ProgressBar({ title, current, target, percentage }: ProgressBarProps) {
    return (
        <div className="bg-white rounded-xl p-6 border border-border shadow-sm">
            <div className="flex items-center justify-between mb-3">
                <h3 className="font-semibold text-foreground">{title}</h3>
                <span className="text-sm font-medium text-primary">{percentage}%</span>
            </div>
            <div className="bg-muted rounded-full h-3 mb-3 overflow-hidden">
                <div
                    className="h-full bg-gradient-to-r from-primary to-primary/80 transition-all duration-300"
                    style={{ width: `${percentage}%` }}
                />
            </div>
            <p className="text-sm text-muted-foreground">
                {current.toLocaleString()} 円 / {target.toLocaleString()} 円
            </p>
        </div>
    )
}
