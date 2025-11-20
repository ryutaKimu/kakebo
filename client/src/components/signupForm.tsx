import { useState } from 'react'
import { User, Mail, Lock } from 'lucide-react'
import { Link, useNavigate } from 'react-router-dom'
import { createAccount } from '@/api/kakebo'
import { ToastError } from './toastNotification'
import { handleApiError } from '@/frontUtils/handleApiError'

export function SignupForm() {
  const navigate = useNavigate()
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [isLoading, setIsLoading] = useState(false)
  const [errorMessage, setErrorMessage] = useState('')



  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setIsLoading(true)
    try {
      await createAccount(name, email, password)
      navigate('/dashboard')
    } catch (err) {
      handleApiError(err, "登録に失敗しました。再度お試しください。", setErrorMessage);
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <>
      {errorMessage && (<ToastError message={errorMessage} onClose={() => setErrorMessage("")} />)}
      <div className="space-y-6">
        <div>
          <h2 className="text-2xl font-bold text-foreground">新規登録</h2>
          <p className="text-muted-foreground mt-1">アカウントを作成してはじめましょう</p>
        </div>

        <form onSubmit={handleSubmit} className="space-y-4">
          {/* 名前入力 */}
          <div>
            <label className="block text-sm font-medium text-foreground mb-2">
              名前
            </label>
            <div className="relative">
              <User className="absolute left-3 top-3 w-5 h-5 text-muted-foreground" />
              <input
                type="text"
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder="山田太郎"
                className="w-full pl-10 pr-4 py-2 border border-border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary bg-background"
                required
              />
            </div>
          </div>

          {/* メールアドレス入力 */}
          <div>
            <label className="block text-sm font-medium text-foreground mb-2">
              メールアドレス
            </label>
            <div className="relative">
              <Mail className="absolute left-3 top-3 w-5 h-5 text-muted-foreground" />
              <input
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                placeholder="you@example.com"
                className="w-full pl-10 pr-4 py-2 border border-border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary bg-background"
                required
              />
            </div>
          </div>

          {/* パスワード入力 */}
          <div>
            <label className="block text-sm font-medium text-foreground mb-2">
              パスワード
            </label>
            <div className="relative">
              <Lock className="absolute left-3 top-3 w-5 h-5 text-muted-foreground" />
              <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="••••••••"
                className="w-full pl-10 pr-4 py-2 border border-border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary bg-background"
                required
              />
            </div>
          </div>

          {/* 登録ボタン */}
          <button
            type="submit"
            disabled={isLoading}
            className="w-full py-2 bg-primary text-primary-foreground rounded-lg font-semibold hover:bg-primary/90 transition-colors disabled:opacity-50"
          >
            {isLoading ? '登録中...' : '新規登録'}
          </button>
        </form>

        {/* ログインリンク */}
        <div className="text-center text-sm">
          <p className="text-muted-foreground">
            すでにアカウントをお持ちですか？{' '}
            <Link to="/login" className="text-primary font-semibold hover:underline">
              ログイン
            </Link>
          </p>
        </div>
      </div>
    </>
  )
}
