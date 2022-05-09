import GitHub from "../../middleware/GitHub";
import { Request, Response } from "express";
import { validateBlackList } from "../../global";

const pathReg = new RegExp("^/gh(/[ut]/\\d+)");

const GitHubAvatar = (req: Request, res: Response) => {
    const path = (req.path.match(pathReg) as string[])[1];
    const size =
        typeof req.query.s !== "undefined" && Number(req.query.s) <= 460
            ? Number(req.query.s)
            : 460;

    if (!validateBlackList("github", path.split("/")[2])) {
        res.sendStatus(403);
        res.end();
        return;
    }

    GitHub.Avatar(path, {
        params: {
            s: size,
            v: 4,
        },
        responseType: "arraybuffer",
    })
        .then((r) => {
            if (r.status === 200) {
                res.statusCode = 200;
                res.contentType("image/jpeg");
                res.setHeader(
                    "Cache-Control",
                    "public, max-age=604800, must-revalidate"
                );
                res.send(Buffer.from(r.data, "binary"));
                res.end();
            }
        })
        .catch((err) => {
            console.error("[Axios]", err.message);
            res.setHeader("Cache-Control", "no-cache");
            res.sendStatus(404);
            res.end();
        })
        .then(() => {
            console.log(
                "[GitHub Avatar]",
                req.method,
                req.originalUrl,
                res.statusCode
            );
        });
};

export default GitHubAvatar;
