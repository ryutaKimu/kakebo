import axios from "axios";

export const handleApiError = (
    err: unknown,
    defaultMessage: string,
    setErrorMessage: (msg: string) => void
) => {
    if (axios.isAxiosError(err)) {
        const data = err.response?.data;
        const apiMessage = typeof data === 'string' ? data : data?.message;

        if (typeof apiMessage === 'string' && apiMessage) {
            setErrorMessage(apiMessage);
        } else {
            console.error("API error with unhandled format:", err);
            setErrorMessage(defaultMessage);
        }
    } else {
        console.error(defaultMessage, err);
        setErrorMessage(defaultMessage);
    }
}