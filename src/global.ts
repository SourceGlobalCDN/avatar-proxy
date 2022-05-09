import fs from "fs";

export interface BlackListType {
    gravatar: string[];
    github: string[];
    githubUser: string[];
}

export const blackList: BlackListType = JSON.parse(
    fs.readFileSync("./blacklist.json").toString("utf8")
);

export const validateBlackList = (type: keyof BlackListType, value: string) => {
    const list = blackList[type];
    for (let item of list) {
        if (item === value) {
            return false;
        }
    }
    return true;
};
