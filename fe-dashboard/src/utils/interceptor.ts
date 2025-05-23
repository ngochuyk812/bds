import type { Interceptor } from "@connectrpc/connect";
import { useAuthStore } from "../store/auth";

export const authInterceptor: Interceptor = (next) => async (req) => {
    const accessToken = localStorage.getItem("auth_token");
    if (accessToken) {
        req.header.set("Authorization", `Bearer ${accessToken}`);
    }
    try {
        const res = await next(req);
        return res;
    } catch (err: any) {
        //401 (UNAUTHENTICATED)
        if (err.code == "16") {
            try {
                const success = await useAuthStore.getState().refreshLogin();
                if (!success) {
                    return Promise.reject(new Error("Token expired. Redirecting."));
                }

                const newToken = localStorage.getItem("auth_token");
                if (newToken) {
                    req.header.set("Authorization", `Bearer ${newToken}`);
                }

                const retryRes = await next(req);
                console.log("response after retry", retryRes);
                return retryRes;

            } catch (refreshErr) {
                console.error("Refresh token failed", refreshErr);
                window.location.href = "/login";
                return Promise.reject(refreshErr);
            }
        }
        throw err;
    }
};
