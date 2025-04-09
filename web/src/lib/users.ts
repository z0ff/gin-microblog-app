import type { User } from './types';
import { PUBLIC_API_ORIGIN } from "$env/static/public";

export const getUser = async (username: string): Promise<User | null> => {
    const res = await fetch(`${PUBLIC_API_ORIGIN}/user/${username}`, {
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

export const searchUser = async (username: string): Promise<User | null> => {
    const res = await fetch(`${PUBLIC_API_ORIGIN}/searchuser?query=${username}`, {
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

export const follow = async (username: string): Promise<boolean> => {
    const res = await fetch(`${PUBLIC_API_ORIGIN}/follow/${username}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        // fetch APIはデフォルトでセッション情報を送信しないため、
        // 明示的に送信するように設定する
        credentials: "include"
    }).catch(console.error);

    if (res) {
        console.log(await res.json());
        return true;
    } else {
        console.error(res);
        return false;
    }
}

export const unfollow = async (username: string): Promise<boolean> => {
    const res = await fetch(`${PUBLIC_API_ORIGIN}/unfollow/${username}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        // fetch APIはデフォルトでセッション情報を送信しないため、
        // 明示的に送信するように設定する
        credentials: "include"
    }).catch(console.error);

    if (res) {
        console.log(await res.json());
        return true;
    } else {
        console.error(res);
        return false;
    }
}

export const isFollowing = async (username: string): Promise<boolean> => {
    const res = await fetch(`${PUBLIC_API_ORIGIN}/is_following/${username}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
        // fetch APIはデフォルトでセッション情報を送信しないため、
        // 明示的に送信するように設定する
        credentials: "include"
    }).catch((error) => {
        throw error;
    });

    if (res) {
        if (res.ok) {
            const data = await res.json();
            return data.is_following;
        } else {
            console.error(res.statusText)
            return false;
        }
    }
}
