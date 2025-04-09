import type {User} from "$lib/types";
import { PUBLIC_API_ORIGIN } from "$env/static/public";

export const getUsers = async (name: string|null = null): Promise<User[]> => {
    const url = PUBLIC_API_ORIGIN + "/admin/users";

    const res = await fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
        credentials: "include",
    }).catch(console.error);

    if (res) {
        const data = await res.json();
        console.log(data);
        if (res.ok) {
            return data;
        } else {
            console.error(data);
            return [];
        }
    } else {
        console.error(res);
        return [];
    }
};

export const getUserDetails = async (id: number): Promise<User | null> => {
    const url = `${PUBLIC_API_ORIGIN}/admin/user/${id}`;

    const res = await fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
        credentials: "include",
    }).catch(console.error);

    if (res) {
        const data = await res.json();
        console.log(data[0]);
        if (res.ok) {
            return data[0];
        } else {
            console.error(data);
            return null;
        }
    } else {
        console.error(res);
        return null;
    }
}

export const deleteUser = async (id: number) => {
    if (confirm("本当に削除しますか？")) {
        await fetch(`${PUBLIC_API_ORIGIN}/admin/user/${id}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: "include"
        }).then(res => {
            if (res.status === 200) {
                location.reload();
            } else {
                alert("削除に失敗しました");
            }
        }).catch(console.error);
    }
};
export const suspendUser = async (id: number) => {
    if (confirm("本当に停止しますか？")) {
        await fetch(`${PUBLIC_API_ORIGIN}/admin/suspend/${id}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: "include"
        }).then(res => {
            if (res.status === 200) {
                location.reload();
            } else {
                alert("停止に失敗しました");
            }
        }).catch(console.error);
    }
};

export const resumeUser = async (id: number) => {
    await fetch(`${PUBLIC_API_ORIGIN}/admin/resume/${id}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        credentials: "include"
    }).then(res => {
        if (res.status === 200) {
            location.reload();
        } else {
            alert("停止に失敗しました");
        }
    }).catch(console.error);
};
