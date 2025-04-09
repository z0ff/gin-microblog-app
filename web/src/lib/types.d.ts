export type User = {
    id: number;
    name: string;
    display_name: string;
    email: string;
    is_followed: boolean;
    // createdAt: Date;
    // updatedAt: Date;
}

export type Post = {
    id: number;
    user: User;
    content: string;
    createdAt: Date | undefined;
    updatedAt: Date | undefined;
    is_liked: boolean | undefined;
}
