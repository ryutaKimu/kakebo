"use client";

import { useState } from "react";
import { Mail, Lock } from "lucide-react";
import { Link } from "react-router-dom";
import { login } from "../api/kakebo";

export function LoginForm() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    try {
      // ここで実際のログイン処理を行います
      console.log(email);

      const data = await login(email, password);

      if (data.token) {
        localStorage.SetItem("access_token", data.token);
      }
    } catch (err) {
      console.error("ログイン失敗:", err);
      alert("ログインに失敗しました");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="space-y-6">
      <div>
        <h2 className="text-2xl font-bold text-foreground">ログイン</h2>
        <p className="text-muted-foreground mt-1">
          アカウントにログインしてください
        </p>
      </div>

      <form onSubmit={handleSubmit} className="space-y-4">
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

        {/* ログインボタン */}
        <button
          type="submit"
          disabled={isLoading}
          className="w-full py-2 bg-primary text-primary-foreground rounded-lg font-semibold hover:bg-primary/90 transition-colors disabled:opacity-50"
        >
          {isLoading ? "ログイン中..." : "ログイン"}
        </button>
      </form>

      {/* 登録リンク */}
      <div className="text-center text-sm">
        <p className="text-muted-foreground">
          アカウントをお持ちでないですか？{" "}
          <Link
            to="/signup"
            className="text-primary font-semibold hover:underline"
          >
            新規登録
          </Link>
        </p>
      </div>
    </div>
  );
}
