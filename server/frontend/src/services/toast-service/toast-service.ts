import { toast, ToastOptions } from "react-toastify";

const options: ToastOptions = {
}

export function toastSuccess(msg: string) {
  toast.success(msg);
}

export function toastFailure(msg: string) {
  toast.error(msg);
}
