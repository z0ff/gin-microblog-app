import type { User } from './types';

export const login = async (email: string, password: string): Promise<boolean> => {
    const res = await fetch('http://localhost:3000/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email: email, password: password }),
        // fetch APIはデフォルトでセッション情報を送信しないため、
        // 明示的に送信するように設定する
        credentials: "include"
    }).catch(console.error);

    if (res) {
        console.log(await res.json());
        return res.ok;
    } else {
        console.error(res);
        return false;
    }
}

export const logout = async (): Promise<boolean> => {
    const res = await fetch('http://localhost:3000/logout', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        credentials: "include"
    }).catch(console.error);

    if (res) {
        console.log(await res.json());
        return res.ok;
    } else {
        console.error(res);
        return false;
    }
}

export const signup = async (name: string, displayName: string, email: string, password: string): Promise<boolean> => {
    const res = await fetch('http://localhost:3000/signup', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            name: name,
            display_name: displayName,
            email: email,
            password: password
        }),
    }).catch(console.error);

    if (res) {
        console.log(await res.json());
        return res.ok;
    } else {
        console.error(res);
        return false;
    }
};

export const getMe = async (): Promise<User | null> => {
    const res = await fetch('http://localhost:3000/me', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
        credentials: "include"
    }).catch(console.error);

    if (res) {
        const data = await res.json();
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
