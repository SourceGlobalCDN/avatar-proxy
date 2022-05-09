import { NextFunction, Request, Response } from "express";

/**
 * @param {*} req
 * @param {*} res
 * @param {NextFunction|Response<*, Record<string, *>>} next
 */
const Security = (req: Request, res: Response, next: NextFunction) => {
    console.log("[Security]", req.method, req.originalUrl, "Time:", Date());
    // UA verification
    if (
        typeof req.headers["user-agent"] === "undefined" ||
        req.headers["user-agent"].length < 5
    ) {
        console.log(
            "[Security]",
            req.method,
            req.originalUrl,
            "Blocked because the UA does not comply with the rules.",
            `(User-Agent: ${req.headers["user-agent"]})`
        );
        res.sendStatus(403);
        res.end();
        return;
    }
    // Query CC attack detection
    for (let paramsKey in req.query) {
        if (paramsKey.length > 20) {
            console.log(
                "[Security]",
                req.method,
                req.originalUrl,
                "Intercepted due to Param exception."
            );
            res.sendStatus(403);
            res.end();
            return;
        }
    }

    // Clear
    next();
};

export default Security;
