"use strict";

/**
 * @param {*} req
 * @param {*} res
 * @param {NextFunction|Response<*, Record<string, *>>} next
 */
exports.Security = (req, res, next) => {
    console.log("[Security]", req.method, req.originalUrl, "Time:", Date.now());
    // UA verification
    if (
        !req.headers.hasOwnProperty("user-agent") ||
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
    if (req.query.length < 2) {
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
    }

    // Clear
    next();
};
