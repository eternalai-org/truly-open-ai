// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

library RequestBuilder {
    function buildRequest(
        string memory _prompt,
        string memory _modelName
    ) internal pure returns (string memory) {
        string memory request = string.concat('{  "messages"');
        request = string.concat(request, ' :[');
        request = string.concat(
            request,
            '{"role":"system","content":"You are a helpful assistant"},'
        );
        request = string.concat(request, '{"role":"user","content":"');
        request = string.concat(request, _prompt);
        request = string.concat(request, ' "}');
        request = string.concat(request, '],');
        request = string.concat(request, '"max_tokens":1024,');
        request = string.concat(request, '"model":"');
        request = string.concat(request, _modelName);
        request = string.concat(request, '"}');
        return request;
    }
}
