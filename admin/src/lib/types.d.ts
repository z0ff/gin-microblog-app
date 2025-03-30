export type User = {
    id: number;
    name: string;
    display_name: string;
    email: string;
    posts: Post[] | null;
    createdAt: Date;
    updatedAt: Date;
    deletedAt: Date | null;
}

export type Post = {
    id: number;
    user_id: number;
    content: string;
    createdAt: Date | undefined;
    updatedAt: Date | undefined;
    deletedAt: Date | undefined;
}

export type Admin = {
    id: number;
    name: string;
    email: string;
    createdAt: Date;
    updatedAt: Date;
    deletedAt: Date | null;
}
