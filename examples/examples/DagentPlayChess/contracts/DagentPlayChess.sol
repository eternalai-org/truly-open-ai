// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

interface AIKernel {
    event NewInference(
        uint256 indexed inferenceId,
        address indexed model,
        address indexed creator,
        uint256 value,
        uint256 originInferenceId
    );

    function infer(
        bytes calldata _data,
        bool _flag
    ) external payable returns (uint256 referenceId);

    function infer(
        bytes calldata _data
    ) external payable returns (uint256 referenceId);
}

interface PromptScheduler {
    enum InferenceStatus {
        Nil,
        Solving,
        Commit,
        Reveal,
        Processed,
        Killed,
        Transferred
    }

    enum AssignmentRole {
        Nil,
        Validating,
        Mining
    }

    enum Vote {
        Nil,
        Disapproval,
        Approval
    }

    struct Assignment {
        uint256 inferenceId;
        bytes32 commitment;
        bytes32 digest;
        uint40 revealNonce;
        address worker;
        AssignmentRole role;
        Vote vote;
        bytes output;
    }

    struct Inference {
        uint256[] assignments;
        bytes input;
        uint256 value;
        uint256 feeL2;
        uint256 feeTreasury;
        address modelAddress;
        uint40 submitTimeout;
        uint40 commitTimeout;
        uint40 revealTimeout;
        InferenceStatus status;
        address creator;
        address processedMiner;
        address referrer;
    }

    function getInferenceInfo(
        uint256 _inferenceId
    ) external view returns (Inference memory);

    function getAssignmentInfo(
        uint256 _assignmentId
    ) external view returns (Assignment memory);
}

contract DagentPlayChess {
    bytes32 private constant INVALID_SIG =
        keccak256(abi.encodePacked("invalid"));
    bytes32 private constant OK_SIG = keccak256(abi.encodePacked("ok"));
    bytes32 private constant EMPTY_SIG = keccak256(abi.encodePacked(""));

    address public kernel;
    address public promptScheduler;
    string public initPlayingPrompt;
    mapping(address => uint256) public currentInferId;
    mapping(address => string) public playingContext;

    event GameCreated(
        address indexed user,
        uint256 inferenceId,
        string request
    );
    event GamePlayed(address indexed user, uint256 inferenceId, string request);

    constructor(address _kernelAddress, address _promptSchedulerAddress) {
        require(
            _kernelAddress != address(0) &&
                _promptSchedulerAddress != address(0),
            "AIPoweredWallet: Invalid address"
        );

        kernel = _kernelAddress;
        promptScheduler = _promptSchedulerAddress;

        string
            memory prompt = "Let's play a text-based chess game.Game Rules:Standard Chess Rules: You will follow standard chess rules, including movement, capturing, check, checkmate, stalemate, castling (both kingside and queenside), en passant, and pawn promotion.Algebraic Notation: Use standard algebraic notation (e.g., e2-e4, Nb1-c3, etc.) for all moves. Turn-Based: You and user will take turns, with White (user) moving first.";
        prompt = string.concat(
            prompt,
            "Use 'e' for the pawn in front of the king, 'd' for the pawn in front of the queen, and so on.For pieces, use the first letter of the piece's name (e.g., N for Knight, B for Bishop, etc.).If you want to specify a rank or file, use the corresponding letter or number (e.g., e2-e4 means the pawn moves from e2 to e4), 'Nb1-c3' to move the knight from b1 to c3.."
        );
        prompt = string.concat(
            prompt,
            "   a  b  c  d  e  f  g  h\\n8  r  n  b  q  k  b  n  r\\n7  p  p  p  p  p  p  p  p\\n6  . .  . .  . .  . .\\n5  . .  . . .  . .  .\\n4  . .  . .  . .  . .\\n3  . .  . .  . .  . .\\n2  P  P  P  P  P  P  P  P\\n1  R  N  B  Q  K  B  N  R"
        );
        initPlayingPrompt = string.concat(
            prompt,
            "Lowercase letters (r, n, b, q, k, p) represent black pieces.\\nUppercase letters (R, N, B, Q, K, P) represent white pieces.\\n '.' represents an empty square.\\n\\nPiece Notation:\\n r/R: Rook\\nn/N: Knight\\nb/B: Bishop\\nq/Q: Queen\\nk/K: King\\np/P: Pawn\\n\\nGameplay Logic and AI:\\n1. Start of Game: If the playing history is empty, respond with ONLY 'ok' to indicate the game has started.\\n2. Move Validation: After each of my moves, validate it against legal chess moves.\\n- If my move is valid: make your move (as Black), ONLY respond with your move like standard algebraic notation (e.g., e2-e4, Nb1-c3, etc.)\\n- If my latest move is invalid: ONLY respond with 'invalid'. Ignore my invalid move for your subsequent moves, waiting for a valid move from user.\\n\\nInput / Output Format:\\n- My Input (White's Move): A single move in algebraic notation (e.g., 'e2-e4', 'Nb1-c3', etc.)  .\\n\\nYour Output (Black's Response):\\n- 'ok' (at the start of the game).\\n- 'invalid' if my latest move is invalid.\\n- Black's move in algebraic notation. "
        );
    }

    function createGame() external {
        string memory request = buildRequest(initPlayingPrompt);
        playingContext[msg.sender] = initPlayingPrompt;

        uint256 inferenceId = AIKernel(kernel).infer(bytes(request));
        currentInferId[msg.sender] = inferenceId;

        emit GameCreated(msg.sender, inferenceId, request);
    }

    function play(string memory x) external {
        uint256 inferId = currentInferId[msg.sender];
        require(inferId != 0, "PlayingChessWithDagent: Invalid id");

        bytes memory result = fetchInferenceResult(inferId);
        bytes32 digest = keccak256(result);

        require(digest != EMPTY_SIG, "PlayingChessWithDagent: Invalid move");

        if (digest == OK_SIG) {
            playingContext[msg.sender] = string.concat(
                playingContext[msg.sender],
                "You: okay. Here is history: "
            );
        } else if (digest == INVALID_SIG) {
            playingContext[msg.sender] = string.concat(
                playingContext[msg.sender],
                "You said: invalid move. "
            );
        } else {
            playingContext[msg.sender] = string.concat(
                playingContext[msg.sender],
                "You played: ",
                string(result),
                ". "
            );
        }

        playingContext[msg.sender] = string.concat(
            playingContext[msg.sender],
            "User played: ",
            x,
            ". "
        );

        string memory request = buildRequest(playingContext[msg.sender]);

        uint256 newInferenceId = AIKernel(kernel).infer(bytes(request));
        currentInferId[msg.sender] = newInferenceId;

        emit GamePlayed(msg.sender, newInferenceId, request);
    }

    function clearPlayingContext() external {
        playingContext[msg.sender] = "";
    }

    function buildRequest(
        string memory _prompt
    ) public pure returns (string memory) {
        string memory request = string.concat('{  "messages"');
        request = string.concat(request, " :[");
        request = string.concat(
            request,
            '{"role":"system","content":"You are a helpful assistant"},'
        );
        request = string.concat(request, '{"role":"user","content":"');
        request = string.concat(request, _prompt);
        request = string.concat(request, ' "}');
        request = string.concat(request, "],");
        request = string.concat(request, '"max_tokens":1024,');
        request = string.concat(
            request,
            '"model":"NousResearch/Hermes-3-Llama-3.1-70B-FP8"'
        );

        request = string.concat(request, "}");
        return request;
    }

    function fetchInferenceResult(
        uint256 _inferenceId
    ) public view returns (bytes memory) {
        PromptScheduler.Inference memory inferInfo = PromptScheduler(
            promptScheduler
        ).getInferenceInfo(_inferenceId);

        if (inferInfo.assignments.length == 0) revert("Wait for inference");

        return
            PromptScheduler(promptScheduler)
                .getAssignmentInfo(inferInfo.assignments[0])
                .output;
    }
}
