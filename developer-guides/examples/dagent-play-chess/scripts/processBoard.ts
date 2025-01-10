type Piece =
  | "r"
  | "n"
  | "b"
  | "q"
  | "k"
  | "p"
  | "R"
  | "N"
  | "B"
  | "Q"
  | "K"
  | "P"
  | ".";

export type Board = Piece[][];

type Move = string; // e.g., 'e4', 'Nf3', 'Qh5'

// Function to print the board
export function printBoard(board: Board) {
  console.log("\n");
  console.log("   a  b  c  d  e  f  g  h");
  for (let row = 7; row >= 0; row--) {
    let rowString = `${row + 1} `;
    for (let col = 0; col < 8; col++) {
      rowString += `${board[row][col]}  `;
    }
    console.log(rowString);
  }
  console.log("\n");
}

function validateMove(move: Move, board: Board, isWhiteTurn: boolean): boolean {
  // Enhanced validation, can be improved

  // Check move format
  const moveRegex =
    /^([a-h][1-8])?([NBRQK])?([a-h][1-8])([a-h][1-8])?([qrbn])?$/;
  const match = move.match(moveRegex);
  if (!match) {
    console.log("Invalid move format.");
    return false;
  }

  const [, startPos, piece, endPos, promotion] = match;

  const [endFile, endRank] = endPos.split("");
  const endIndex = [
    parseInt(endRank) - 1,
    endFile.charCodeAt(0) - "a".charCodeAt(0),
  ];

  let startFile: string | undefined;
  let startRank: string | undefined;
  let startIndex: number[] = [];

  if (startPos) {
    const [f, r] = startPos.split("");

    if (r && f) {
      startFile = f;
      startRank = r;
      startIndex = [
        parseInt(startRank) - 1,
        startFile.charCodeAt(0) - "a".charCodeAt(0),
      ];
    }
  } else {
    if (!piece) {
      console.log("Invalid move: no starting pos found");
      return false;
    }
  }

  if (startIndex.length > 0 && board[startIndex[0]][startIndex[1]] === ".") {
    console.log("Invalid move: no piece at starting position");
    return false;
  }

  let startPiece: Piece;
  if (startIndex.length > 0) {
    startPiece = board[startIndex[0]][startIndex[1]];
    if (
      isWhiteTurn &&
      startPiece.toLowerCase() === startPiece &&
      !isWhiteTurn &&
      startPiece.toUpperCase() === startPiece
    ) {
      console.log("Invalid move: wrong turn");
      return false;
    }
  } else {
    if (piece) {
      let found = false;
      for (let i = 0; i < 8; i++) {
        for (let j = 0; j < 8; j++) {
          const currentPiece = board[i][j];

          if (
            (isWhiteTurn && currentPiece === piece) ||
            (!isWhiteTurn && currentPiece.toLowerCase() === piece.toLowerCase())
          ) {
            startIndex = [i, j];
            startPiece = currentPiece;
            found = true;
            break;
          }
        }
        if (found) {
          break;
        }
      }

      if (!found) {
        console.log("Invalid move: no piece to move");
        return false;
      }
    }
  }

  return true;
}

function updateBoard(move: Move, board: Board, isWhiteTurn: boolean): Board {
  console.log("Updating board...");
  // basic move implementation (can be expanded based on piece)
  const moveRegex =
    /^([a-h][1-8])?([NBRQK])?([a-h][1-8])([a-h][1-8])?([qrbn])?$/;
  const match = move.match(moveRegex);

  if (!match) {
    return board;
  }

  const [, , piece, ,] = match;

  let match1;
  if (move.length == 4) {
    match1 = match;
  } else {
    match1 = move.slice(1).match(moveRegex);
  }

  if (!match1) {
    console.log("Invalid move format.");
    return board;
  }

  const [, startPos, , endPos] = match1;

  const [endFile, endRank] = endPos.split("");
  const endIndex = [
    parseInt(endRank) - 1,
    endFile.charCodeAt(0) - "a".charCodeAt(0),
  ];

  let startFile: string | undefined;
  let startRank: string | undefined;
  let startIndex: number[] = [];
  let startPiece: Piece = currentBoard[0][0]; // Declare startPiece here

  if (startPos) {
    const [f, r] = startPos.split("");

    if (r && f) {
      startFile = f;
      startRank = r;
      startIndex = [
        parseInt(startRank) - 1,
        startFile.charCodeAt(0) - "a".charCodeAt(0),
      ];
      startPiece = board[startIndex[0]][startIndex[1]];
    }
  } else {
    if (piece) {
      let found = false;
      for (let i = 0; i < 8; i++) {
        for (let j = 0; j < 8; j++) {
          const currentPiece = board[i][j];
          if (
            (isWhiteTurn && currentPiece == piece.toLowerCase()) ||
            (!isWhiteTurn && currentPiece.toLowerCase() === piece.toLowerCase())
          ) {
            startIndex = [i, j];
            startPiece = currentPiece; //Assign start piece here
            found = true;
            break;
          }
        }
        if (found) {
          break;
        }
      }
    }
  }

  const updatedBoard = board.map((row) => [...row]);
  if (startIndex.length > 0) {
    updatedBoard[endIndex[0]][endIndex[1]] = startPiece;
    updatedBoard[startIndex[0]][startIndex[1]] = ".";
  }

  return updatedBoard;
}

function processMove(move: Move, board: Board, isWhiteTurn: boolean): Board {
  move = removeDashes(move);

  if (!validateMove(move, board, isWhiteTurn)) {
    return board;
  }

  const updatedBoard = updateBoard(move, board, isWhiteTurn);
  printBoard(updatedBoard);
  return updatedBoard;
}

function removeDashes(str: string) {
  return str.replace("-", "");
}

// Initialize the board
export let currentBoard: Board = [
  ["r", "n", "b", "q", "k", "b", "n", "r"],
  ["p", "p", "p", "p", "p", "p", "p", "p"],
  [".", ".", ".", ".", ".", ".", ".", "."],
  [".", ".", ".", ".", ".", ".", ".", "."],
  [".", ".", ".", ".", ".", ".", ".", "."],
  [".", ".", ".", ".", ".", ".", ".", "."],
  ["P", "P", "P", "P", "P", "P", "P", "P"],
  ["R", "N", "B", "Q", "K", "B", "N", "R"],
];

// Game Loop (Example)
let whiteTurn = true;

export function gameMove(move: string) {
  currentBoard = processMove(move.trim(), currentBoard, whiteTurn);
  whiteTurn = !whiteTurn;
}

// Example usage:
// printBoard(currentBoard);
// gameLoop("b2-b4");
// gameLoop("c7-c5");
// gameLoop("Nb1-c3");
// gameLoop("e2-e4");
// gameLoop("Bb5");
