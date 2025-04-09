import type { Post } from '$lib/types';
import { PUBLIC_API_ORIGIN } from "$env/static/public";

export const getPosts = async (): Promise<Post[]> => {
    const url = PUBLIC_API_ORIGIN + "/admin/posts";

    const res = await fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
        credentials: "include",
    }).catch(console.error);

    if (res) {
        const data: Post[] = await res.json();
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

export const deletePost = async (id: number) => {
    if (confirm("本当に削除しますか？")) {
        await fetch(`${PUBLIC_API_ORIGIN}/admin/post/${id}`, {
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