import { toast, ToastOptions } from "react-toastify";

const options: ToastOptions = {
  autoClose: 2000,
}

export function toastSuccess(msg: string) {
  toast.success(msg, options);
}

export function toastFailure(msg: string) {
  toast.error(msg, options);
}
