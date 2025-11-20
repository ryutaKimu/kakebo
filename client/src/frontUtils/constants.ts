export const VALIDATION = {
PASSWORD_MIN_LENGTH: 8,
EMPTY_NAME: "名前を入力してください。",
EMPTY_EMAIL: "メールアドレスを入力してください。",
EMPTY_PASSWORD: "パスワードを入力してください。",
SHORT_PASSWORD: (min: number) =>
`パスワードは${min}文字以上で入力してください。`,
};

export const API_ERROR = {
SIGNUP: "登録に失敗しました。再度お試しください。",
LOGIN: "ログインに失敗しました。再度お試しください。",
FETCH_DATA: "データの取得に失敗しました。時間をおいて再度お試しください。",
};
