import axios from "axios";

export const handleApiError = (
    err: unknown,
    defaultMessage: string,
    setErrorMessage: (msg: string) => void
) => {
    if (axios.isAxiosError(err)) {
        // 4xx系のクライアントエラーの場合、バリデーションフィードバックなど
        // 有用な情報が含まれていることが多いため、APIメッセージを表示します。
        if (err.response && err.response.status >= 400 && err.response.status < 500) {
            const data = err.response.data;
            const apiMessage = typeof data === 'string' ? data : data?.message;

            if (typeof apiMessage === 'string' && apiMessage) {
                console.error("API error:", err);
                setErrorMessage(apiMessage);
                return;
            }
        }

        // 5xx系のサーバーエラーやその他の予期せぬエラーの場合は、
        // 詳細をログに出力し、汎用的なメッセージを表示します。
        console.error("API error:", err);
        setErrorMessage(defaultMessage);
    } else {
        console.error(defaultMessage, err);
        setErrorMessage(defaultMessage);
    }
}
