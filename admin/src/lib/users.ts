import type {User} from "$lib/types";

export const getUsers = async (name: string|null = null): Promise<User[]> => {
    let url = "http://localhost:3000/admin/users";

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
    let url = `http://localhost:3000/admin/user/${id}`;

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
        await fetch(`http://localhost:3000/admin/user/${id}`, {
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
