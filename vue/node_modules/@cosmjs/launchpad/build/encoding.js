"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.makeSignBytes = void 0;
/* eslint-disable @typescript-eslint/naming-convention */
const encoding_1 = require("@cosmjs/encoding");
const lcdapi_1 = require("./lcdapi");
function sortJson(json) {
    if (typeof json !== "object" || json === null) {
        return json;
    }
    if (Array.isArray(json)) {
        return json.map(sortJson);
    }
    const sortedKeys = Object.keys(json).sort();
    const result = sortedKeys.reduce((accumulator, key) => (Object.assign(Object.assign({}, accumulator), { [key]: sortJson(json[key]) })), {});
    return result;
}
function makeSignBytes(msgs, fee, chainId, memo, accountNumber, sequence) {
    const signDoc = {
        account_number: lcdapi_1.uint64ToString(accountNumber),
        chain_id: chainId,
        fee: fee,
        memo: memo,
        msgs: msgs,
        sequence: lcdapi_1.uint64ToString(sequence),
    };
    const sortedSignDoc = sortJson(signDoc);
    return encoding_1.toUtf8(JSON.stringify(sortedSignDoc));
}
exports.makeSignBytes = makeSignBytes;
//# sourceMappingURL=encoding.js.map