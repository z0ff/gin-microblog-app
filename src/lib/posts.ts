import { PUBLIC_API_ORIGIN } from "$env/static/public";
import type { Post } from "$lib/types";

export const getTimeline = async (): Promise<Post[] | null> => {
    const res = await fetch(`${PUBLIC_API_ORIGIN}/`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
        // fetch APIはデフォルトでセッション情報を送信しないため、
        // 明示的に送信するように設定する
        credentials: "include"
    }).catch(console.error);

    if (res) {
        return await res.json() as Post[];
    } else {
        console.error(res);
        return null;
    }
};

export const submitPost = async (content: string): Promise<boolean> => {
    const res = await fetch(`${PUBLIC_API_ORIGIN}/post`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        // fetch APIはデフォルトでセッション情報を送信しないため、
        // 明示的に送信するように設定する
        credentials: "include",
        body: JSON.stringify({ content: content }),
    }).catch(console.error);

    if (res) {
        console.log(await res.json());
        return true;
    } else {
        console.error(res);
        return false;
    }
};

export const searchPosts = async (query: string | null):Promise<Post[] | null> => {
    const res = await fetch(`${PUBLIC_API_ORIGIN}/search?query=${query}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
        // fetch APIはデフォルトでセッション情報を送信しないため、
        // 明示的に送信するように設定する
        credentials: "include"
    }).catch(console.error);

    if (res) {
        return await res.json() as Post[];
    } else {
        return null;
    }
};

export const getPostsOfUser = async (username: string): Promise<Post[] | null> => {
    const res = await fetch(`${PUBLIC_API_ORIGIN}/post/${username}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
        // fetch APIはデフォルトでセッション情報を送信しないため、
        // 明示的に送信するように設定する
        credentials: "include"
    }).catch((error) => {
        console.error(error);
        return;
    });

    if (res) {
        return await res.json() as Post[];
    } else {
        console.error(res);
        return null;
    }

};
export const like = async (postID): Promise<boolean> => {
    const res = await fetch(`${PUBLIC_API_ORIGIN}/like/${postID}`, {
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

export const unlike = async (postID): Promise<boolean> => {
    const res = await fetch(`${PUBLIC_API_ORIGIN}/unlike/${postID}`, {
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
