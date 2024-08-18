/**
 * Assume the following rules are for the tic-tac-toe game on an n x n board between two players:
 *
 * A move is guaranteed to be valid and is placed on an empty block.
 * Once a winning condition is reached, no more moves are allowed.
 * A player who succeeds in placing n of their marks in a horizontal, vertical, or diagonal row wins the game.
 * Implement the TicTacToe class:
 *
 * TicTacToe(int n) Initializes the object the size of the board n.
 * int move(int row, int col, int player) Indicates that the player with id player plays at the cell (row, col) of
 * the board. The move is guaranteed to be a valid move, and the two players alternate in making moves. Return
 * 0 if there is no winner after the move,
 * 1 if player 1 is the winner after the move, or
 * 2 if player 2 is the winner after the move.
 *
 * Constraints:
 * 2 <= n <= 100
 * player is 1 or 2.
 * 0 <= row, col < n
 * (row, col) are unique for each different call to move.
 * At most n2 calls will be made to move.
 *
 * ------------------------------------------------------
 * Trick here is understanding you don't need to traverse the board on every move to determine if
 * there is a winner. A big advantage we have is players can only make valid moves so we can use variables
 * to keep track of each row, column, diagonal and anti-diagonal. Identifying a winner is as simple as checking if
 * the current column,row, diagonal, or anti-diagonal matches the board size.
 *
 * Identifying a row or column is each rows[row] and cols[col].
 * Identifying if we are on a diagonal is easy. row === col
 * The hardest part of the problem is identifying the anti-diagonal which can be done with "boardSize - row - 1"
 */

/**
 * @param {number} n
 */
var TicTacToe = function(n) {
  this.boardSize = n
  this.rows = new Array(n).fill(0)
  this.cols = new Array(n).fill(0)
  this.diagonal = 0
  this.antiDiagonal = 0
};

/**
 * @param {number} row
 * @param {number} col
 * @param {number} player
 * @return {number}
 */
TicTacToe.prototype.move = function(row, col, player) {
  const playerCounter = player === 1 ? 1 : -1
  this.rows[row] += playerCounter
  this.cols[col] += playerCounter

  if (col === row) {
    this.diagonal += playerCounter
  }

  if (col === this.boardSize - row - 1) {
    this.antiDiagonal += playerCounter
  }

  if (
    Math.abs(this.rows[row]) === this.boardSize ||
    Math.abs(this.cols[col]) === this.boardSize ||
    Math.abs(this.diagonal) === this.boardSize ||
    Math.abs(this.antiDiagonal) === this.boardSize
  ) return player

  return 0
};

/**
 * Your TicTacToe object will be instantiated and called as such:
 * var obj = new TicTacToe(n)
 * var param_1 = obj.move(row,col,player)
 */
