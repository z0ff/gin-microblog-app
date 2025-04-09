import type { Admin } from "$lib/types";
import { PUBLIC_API_ORIGIN } from "$env/static/public";

export const login = async (email: string, password: string): Promise<Admin | null> => {
    const url = PUBLIC_API_ORIGIN + "/admin/auth/login";

    const res = await fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify({ email, password }),
    }).catch(console.error);

    if (res) {
        const data = await res.json();
        console.log(data);
        if (res.ok) {
            return data;
        } else {
            console.error(data);
            return null;
        }
    } else {
        console.error(res);
        return null;
    }
}

export const logout = async (): Promise<boolean> => {
    const url = PUBLIC_API_ORIGIN + "/admin/auth/logout";

    const res = await fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        credentials: "include",
    }).catch(console.error);

    if (res) {
        const data = await res.json();
        console.log(data);
        if (res.ok) {
            return true;
        } else {
            console.error(data);
            return false;
        }
    } else {
        console.error(res);
        return false;
    }
}

export const signup = async (name: string, email: string, password: string): Promise<Admin | null> => {
    const url = PUBLIC_API_ORIGIN + "/admin/auth/signup";

    const res = await fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify({ name, email, password }),
    }).catch(console.error);

    if (res) {
        const data = await res.json();
        console.log(data);
        if (res.ok) {
            return data;
        } else {
            console.error(data);
            return null;
        }
    } else {
        console.error(res);
        return null;
    }
}

export const isLoggedIn = async (): Promise<Admin | null> => {
    const url = PUBLIC_API_ORIGIN + "/admin/auth/is_logged_in";

    const res = await fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
        credentials: "include",
    }).catch(console.error);

    if (res) {
        const data: Admin = await res.json();
        console.log(data);
        if (res.ok) {
            return data;
        } else {
            console.error(data);
            return null;
        }
    } else {
        console.error(res);
        return null;
    }
}

export const getMe = async (): Promise<Admin | null> => {
    const url = PUBLIC_API_ORIGIN + "/admin/auth/me";

    const res = await fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
        credentials: "include",
    }).catch(console.error);

    if (res) {
        const data: Admin = await res.json();
        console.log(data);
        if (res.ok) {
            return data;
        } else {
            console.error(data);
            return null;
        }
    } else {
        console.error(res);
        return null;
    }
}