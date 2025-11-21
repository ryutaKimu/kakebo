
import { LogOut } from 'lucide-react'
import { Link } from 'react-router-dom'

export function AppHeader() {
  return (
    <header className="sticky top-0 z-40 bg-white border-b border-border">
      <div className="max-w-7xl mx-auto px-6 py-4 flex items-center justify-between">
        <Link to="/dashboard" className="flex items-center gap-2">
          <div className="w-10 h-10 bg-primary rounded-lg flex items-center justify-center">
            <span className="text-white font-bold text-lg">F</span>
          </div>
          <span className="text-xl font-bold text-foreground hidden sm:inline">FinTrack</span>
        </Link>
        <button className="p-2 hover:bg-muted rounded-lg transition-colors">
          <LogOut className="w-5 h-5 text-muted-foreground" />
        </button>
      </div>
    </header>
  )
}
