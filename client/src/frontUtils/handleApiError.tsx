import axios from "axios";

export const handleApiError = (
    err: unknown,
    defaultMessage: string,
    setErrorMessage: (msg: string) => void
) => {
    if (axios.isAxiosError(err)) {
        const message =
            err.response?.data?.message ||
            err.response?.data ||
            defaultMessage;

        setErrorMessage(String(message));
    } else {
        console.error(defaultMessage, err);
        setErrorMessage(defaultMessage);
    }
}