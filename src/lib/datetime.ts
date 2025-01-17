export const strToDate = (dateStr: string) => {
    return new Date(dateStr);
}

export const utcToLocalTime = (utcDate: Date) => {
    return new Date(utcDate.getTime() + (new Date().getTimezoneOffset() * 60 * 1000));
}