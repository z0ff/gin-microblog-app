import type { Post } from '$lib/types';

export const getPosts = async (): Promise<Post[]> => {
    const res = await fetch('https://jsonplaceholder.typicode.com/posts');
    return res.json();
};

export const getPostsOfUser = async (userId: number): Promise<Post[]> => {
    const res = await fetch(`https://jsonplaceholder.typicode.com/posts?userId=${userId}`);
    return res.json();
};