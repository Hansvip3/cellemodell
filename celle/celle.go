package cell

func InitCell() bool { return false }



/*
   INN-DATA --> FUNKSJON --> UT-DATA
      ()        InitCell      etter InitCell()
*/


var cell bool

func initCell() bool {
    cell = false;
    return cell
}


func SetCellValue(value bool) bool {
    cell = value;
    return cell;
}

func GetCellValue() bool {
    return cell
}