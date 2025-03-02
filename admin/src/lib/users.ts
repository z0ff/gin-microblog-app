import type {User} from "$lib/types";

export const getUsers = async (): Promise<User[]> => {
    const res = await fetch("http://localhost:3000/admin/users");
    return res.json();
};