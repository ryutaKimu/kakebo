import { TrendingUp, Target, BarChart3 } from 'lucide-react'
import { Link } from 'react-router-dom'

export default function Home() {
  return (
    <main className="min-h-screen bg-gradient-to-b from-background to-secondary">
      {/* ナビゲーションバー */}
      <nav className="sticky top-0 z-50 bg-white/80 backdrop-blur-sm border-b border-border">
        <div className="max-w-7xl mx-auto px-6 py-4 flex items-center justify-between">
          <div className="text-xl font-bold text-primary">FinTrack</div>
          <Link to="/login" className="px-6 py-2 bg-primary text-primary-foreground rounded-lg font-medium hover:bg-primary/90 transition-colors">
            ログイン
          </Link>
        </div>
      </nav>

      {/* ヒーロー セクション */}
      <section className="max-w-7xl mx-auto px-6 py-20 flex flex-col items-center text-center gap-8">
        <h1 className="text-5xl md:text-6xl font-bold text-balance text-foreground">
          シンプルで <span className="text-primary">清潔感のある</span>
          <br />
          家計簿アプリ
        </h1>
        <p className="text-xl text-muted-foreground max-w-2xl text-balance">
          毎日の収支をサッと記録。貯金目標の進捗も一目でわかる。
          複雑な操作は一切なし。シンプルに、正確に。
        </p>
        <Link to="/signup" className="px-8 py-3 bg-primary text-primary-foreground rounded-lg font-semibold text-lg hover:bg-primary/90 transition-all hover:shadow-lg">
          無料ではじめる
        </Link>
      </section>

      {/* 特徴セクション */}
      <section className="max-w-7xl mx-auto px-6 py-20">
        <h2 className="text-3xl font-bold text-center mb-16 text-foreground">
          FinTrack の特徴
        </h2>
        <div className="grid md:grid-cols-3 gap-8">
          {/* 特徴1 */}
          <div className="bg-white rounded-xl p-8 border border-border shadow-sm hover:shadow-md transition-shadow">
            <div className="mb-4 inline-flex p-3 bg-secondary rounded-lg">
              <TrendingUp className="w-6 h-6 text-primary" />
            </div>
            <h3 className="text-xl font-semibold mb-3 text-foreground">
              かんたん記録
            </h3>
            <p className="text-muted-foreground leading-relaxed">
              日付、金額、カテゴリを選ぶだけで記録完了。毎日3秒でできます。
            </p>
          </div>

          {/* 特徴2 */}
          <div className="bg-white rounded-xl p-8 border border-border shadow-sm hover:shadow-md transition-shadow">
            <div className="mb-4 inline-flex p-3 bg-secondary rounded-lg">
              <BarChart3 className="w-6 h-6 text-primary" />
            </div>
            <h3 className="text-xl font-semibold mb-3 text-foreground">
              可視化
            </h3>
            <p className="text-muted-foreground leading-relaxed">
              グラフやチャートで支出パターンが丸わかり。
              無駄な出費も自動で見つかります。
            </p>
          </div>

          {/* 特徴3 */}
          <div className="bg-white rounded-xl p-8 border border-border shadow-sm hover:shadow-md transition-shadow">
            <div className="mb-4 inline-flex p-3 bg-secondary rounded-lg">
              <Target className="w-6 h-6 text-primary" />
            </div>
            <h3 className="text-xl font-semibold mb-3 text-foreground">
              貯金目標管理
            </h3>
            <p className="text-muted-foreground leading-relaxed">
              目標設定して、進捗をリアルタイム追跡。
              達成まであと○○円の表示でモチベーションアップ。
            </p>
          </div>
        </div>
      </section>

      {/* CTA セクション */}
      <section className="bg-gradient-to-r from-primary to-primary/80 text-primary-foreground py-20">
        <div className="max-w-4xl mx-auto px-6 text-center">
          <h2 className="text-4xl font-bold mb-6 text-balance">
            今すぐ、あなたの家計を見える化
          </h2>
          <p className="text-lg mb-8 opacity-90 text-balance">
            登録も設定も簡単。1分で始められます。
          </p>
          <Link to="/signup" className="inline-block px-8 py-3 bg-white text-primary rounded-lg font-semibold hover:bg-secondary transition-colors">
            無料登録する
          </Link>
        </div>
      </section>

      {/* フッター */}
      <footer className="bg-foreground/5 py-8 border-t border-border">
        <div className="max-w-7xl mx-auto px-6 text-center text-sm text-muted-foreground">
          <p>© 2025 FinTrack. All rights reserved.</p>
        </div>
      </footer>
    </main>
  )
}
