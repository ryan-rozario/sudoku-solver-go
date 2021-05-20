let gen_btn = document.getElementById("generate")
gen_btn.addEventListener("click", createGrid);
var numberOfRows = 0;
var numberOfCols = 0;

function createGrid() {
    var Container = document.getElementsByClassName("grid-container")[0];
    Container.innerHTML = "";
    numberOfRows = document.getElementById("rows_num").value;
    numberOfCols = document.getElementById("cols_num").value;
    document.documentElement.style.setProperty("--columns-no", numberOfCols);
    document.documentElement.style.setProperty("--row-no", numberOfRows);
    console.log(numberOfRows, numberOfCols)
    for (i =  0; i < numberOfRows ; i++) {
        for (j =  0; j < numberOfCols ; j++) {
            var inp_ele = document.createElement("input");
            inp_ele.className = "sudoku-cell";
            inp_ele.id = "item-"+i.toString()+"-"+j.toString();
            inp_ele.minLength = "1";
            inp_ele.pattern = "^[0-9]+$/";
            inp_ele.value = "";
            inp_ele.type = "text" ;
            document.getElementsByClassName("grid-container")[0].appendChild(inp_ele);
        }
    }
}

let solve_btn = document.getElementById("solve")
solve_btn.addEventListener("click", submitData);

function submitData() {
    var cellid;

    var arr = new Array(numberOfRows);

    for (i =  0; i < numberOfRows ; i++) {
        arr[i] = new Array(numberOfCols);
        for (j =  0; j < numberOfCols ; j++) {
            cellid = "item-"+i.toString()+"-"+j.toString();
            ele = document.getElementById(cellid).value
            if(ele==""){
                ele="0";
            }
            arr[i][j] = parseInt(ele)
        }
    }

    console.log(JSON.stringify(arr))
}