package main

import (
	"fmt"
)

func knightMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {
	if (targetRow-sourceRow == 2 || targetRow-sourceRow == -2) && (targetCol-sourceCol == 1 || targetCol-sourceCol == -1) {
		return true
	}
	return false
}

func bishopMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {

	if targetRow < sourceRow && targetCol < sourceCol && sourceRow-targetRow == sourceCol-targetCol { //left up
		for i, j := sourceRow-1, sourceCol-1; i > targetRow; i, j = i-1, j-1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the bishop")
				return false
			}
		}

	} else if targetRow > sourceRow && targetCol < sourceCol && targetRow-sourceRow == sourceCol-targetCol { //left down
		for i, j := sourceRow+1, sourceCol-1; i < targetRow; i, j = i+1, j-1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the bishop")
				return false
			}
		}
	} else if targetRow > sourceRow && targetCol > sourceCol && targetRow-sourceRow == targetCol-sourceCol { //right down
		for i, j := sourceRow+1, sourceCol+1; i < targetRow; i, j = i+1, j+1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the bishop")
				return false
			}
		}
	} else if targetRow < sourceRow && targetCol > sourceCol && sourceRow-targetRow == targetCol-sourceCol { //right up
		for i, j := sourceRow-1, sourceCol+1; i < targetRow; i, j = i-1, j+1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the bishop")
				return false
			}
		}
	} else {
		fmt.Println("Invalid bishop move")
		return false
	}
	return true
}
//combination of bishop and rook
func queenMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {
	
	//check diagonals using bishop moves
	if targetRow < sourceRow && targetCol < sourceCol && sourceRow-targetRow == sourceCol-targetCol { //left up
		for i, j := sourceRow-1, sourceCol-1; i > targetRow; i, j = i-1, j-1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}

	} else if targetRow > sourceRow && targetCol < sourceCol && targetRow-sourceRow == sourceCol-targetCol { //left down
		for i, j := sourceRow+1, sourceCol-1; i < targetRow; i, j = i+1, j-1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}
	} else if targetRow > sourceRow && targetCol > sourceCol && targetRow-sourceRow == targetCol-sourceCol { //right down
		for i, j := sourceRow+1, sourceCol+1; i < targetRow; i, j = i+1, j+1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}
	} else if targetRow < sourceRow && targetCol > sourceCol && sourceRow-targetRow == targetCol-sourceCol { //right up
		for i, j := sourceRow-1, sourceCol+1; i < targetRow; i, j = i-1, j+1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}
	}else if sourceRow<targetRow && sourceCol==targetCol{           //up
		for i := sourceRow+1; i<targetRow; i++{
			if ChessBoard[i][sourceCol] != "--"{
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}
	
	}else if sourceRow > targetRow && sourceCol==targetCol{   //down
		for i := sourceRow-1; i>targetRow; i--{
			if ChessBoard[i][sourceCol] != "--"{
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}
	}else if targetCol>sourceCol && sourceRow==targetRow{     //right
		for i := sourceCol+1; i<targetCol; i++{
			if ChessBoard[sourceRow][i] != "--"{
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}
	}else if targetCol<sourceCol && sourceRow==targetRow{     //left
		for i := sourceCol-1; i>targetCol; i--{
			if ChessBoard[sourceRow][i] != "--"{
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}
	} else{
		fmt.Println("Invalid queen move 2")	
		return false
	}
	return true
}

func rookMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {
	if sourceRow<targetRow && sourceCol==targetCol{           //up
		for i := sourceRow+1; i<targetRow; i++{
			if ChessBoard[i][sourceCol] != "--"{
				fmt.Println("Piece is blocking the rook")
				return false
			}
		}
	
	}else if sourceRow > targetRow && sourceCol==targetCol{   //down
		for i := sourceRow-1; i>targetRow; i--{
			if ChessBoard[i][sourceCol] != "--"{
				fmt.Println("Piece is blocking the rook")
				return false
			}
		}
	}else if targetCol>sourceCol && sourceRow==targetRow{     //right
		for i := sourceCol+1; i<targetCol; i++{
			if ChessBoard[sourceRow][i] != "--"{
				fmt.Println("Piece is blocking the rook")
				return false
			}
		}
	}else if targetCol<sourceCol && sourceRow==targetRow{     //left
		for i := sourceCol-1; i>targetCol; i--{
			if ChessBoard[sourceRow][i] != "--"{
				fmt.Println("Piece is blocking the rook")
				return false
			}
		}
	} else{
		fmt.Println("Invalid rook move")	
		return false
	}
	return true
}
//if king moves two spaces to right as white it checks for kingside castle, three spaces to the left as white checks for white queen castle
func kingMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {
	return true
}
