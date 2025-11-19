export function AuthLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="min-h-screen flex flex-col md:flex-row">
      {/* 左側：ブランドセクション */}
      <div className="hidden md:flex md:w-1/2 bg-gradient-to-br from-primary to-primary/80 text-primary-foreground flex-col justify-between p-12">
        <div>
          <h1 className="text-3xl font-bold">FinTrack</h1>
          <p className="text-primary-foreground/80 mt-2">シンプルな家計簿</p>
        </div>
        <div className="space-y-6">
          <div>
            <h3 className="text-lg font-semibold mb-2">かんたん記録</h3>
            <p className="text-primary-foreground/80">毎日の収支をサッと記録</p>
          </div>
          <div>
            <h3 className="text-lg font-semibold mb-2">一目でわかる</h3>
            <p className="text-primary-foreground/80">グラフで支出パターン可視化</p>
          </div>
          <div>
            <h3 className="text-lg font-semibold mb-2">目標管理</h3>
            <p className="text-primary-foreground/80">貯金目標の進捗をリアルタイム追跡</p>
          </div>
        </div>
      </div>

      {/* 右側：フォームセクション */}
      <div className="w-full md:w-1/2 flex flex-col items-center justify-center p-6">
        <div className="w-full max-w-sm">
          {children}
        </div>
      </div>
    </div>
  )
}
