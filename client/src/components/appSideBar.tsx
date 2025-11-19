import { Home, TrendingUp, TrendingDown, Target } from 'lucide-react'
import { Link, useLocation } from 'react-router-dom'

const navItems = [
  { href: '/dashboard', label: 'ダッシュボード', icon: Home },
  { href: '/income', label: '収入管理', icon: TrendingUp },
  { href: '/expense', label: '支出管理', icon: TrendingDown },
  { href: '/goals', label: '貯金目標', icon: Target },
]

export function AppSidebar() {
  const pathname = useLocation().pathname


  return (
    <aside className="w-64 bg-white border-r border-border h-[calc(100vh-65px)] overflow-y-auto">
      <nav className="p-6 space-y-2">
        {navItems.map((item) => {
          const Icon = item.icon
          const isActive = pathname === item.href
          return (
            <Link
              key={item.href}
              to={item.href}
              className={`flex items-center gap-3 px-4 py-3 rounded-lg transition-colors ${
                isActive
                  ? 'bg-secondary text-primary font-semibold'
                  : 'text-muted-foreground hover:bg-muted'
              }`}
            >
              <Icon className="w-5 h-5" />
              <span>{item.label}</span>
            </Link>
          )
        })}
      </nav>
    </aside>
  )
}
