export type User = {
    id: number;
    name: string;
    display_name: string;
    email: string;
    createdAt: Date;
    updatedAt: Date;
}

export type Post = {
    id: number;
    user: User;
    content: string;
    createdAt: Date | undefined;
    updatedAt: Date | undefined;
}
