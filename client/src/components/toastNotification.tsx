// components/ui/ToastError.tsx
import { X } from "lucide-react";

export const ToastError = ({ message, onClose }: { message: string; onClose: () => void }) => {
    return (
        <div
            id="toast-danger"
            className="flex bg-red-500 text-white items-center w-full max-w-sm p-4 rounded-lg shadow-lg border border-red-700 fixed top-4 right-4 animate-slide-in"
            role="alert"
        >
            <div className="ml-3 text-sm font-normal">{message}</div>

            <button
                onClick={onClose}
                className="ml-auto flex items-center justify-center hover:bg-red-600 rounded h-8 w-8"
            >
                <X className="w-5 h-5" />
            </button>
        </div>
    );
};
