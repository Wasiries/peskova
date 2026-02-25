use std::io;

fn gas(matrix: &mut Vec<Vec<f64>>, vector: &mut Vec<f64>) -> Option<Vec<f64>> {

    // let mut mid: Vec<f64> = Vec::new();
    // let mut right: Vec<f64> = Vec::new();
    // let mut left: Vec<f64> = Vec::new();
    // let mut result = vector.clone();

    // for i in 0..matrix.len() {
    //     if i > 0 {
    //         right.push(matrix[i - 1][i]);
    //         left.push(matrix[i][i - 1]);
    //     }
    //     mid.push(matrix[i][i]);
    // }

    // return Some(result);



    let n = matrix.len();
    let eps = 1e-12;
    for k in 0..n {
        let mut max_row = k;
        let mut max_val = matrix[k][k].abs();
        for i in (k + 1)..n {
            let val = matrix[i][k].abs();
            if val > max_val {
                max_val = val;
                max_row = i;
            }
        }
        if max_val < eps {
            return None;
        }
        matrix.swap(k, max_row);
        vector.swap(k, max_row);
        for i in (k + 1)..n {
            let coef = matrix[i][k] / matrix[k][k];
            for j in k..n {
                matrix[i][j] -= coef * matrix[k][j];
            }
            vector[i] -= coef * vector[k];
        }
    }


    return Some(vector.clone());
}

fn input() -> f64 {
    let mut value: String = String::new();
    std::io::stdin().read_line(&mut value).unwrap();
    let value: f64 = value.trim().parse().unwrap();
    return value;
}

fn line_input() -> Vec<f64> {
    let mut value: String = String::new();
    std::io::stdin().read_line(&mut value).unwrap();
    let value: Vec<f64> = value.trim().split_whitespace().map(|x| x.parse().unwrap()).collect();
    return value;
}


fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut vector: Vec<f64> = vec![7.0, 14.0, 11.0];
    let mut matrix: Vec<Vec<f64>> = vec![
                                vec![2.0, 3.0, -1.0], 
                                vec![1.0, -1.0, 6.0], 
                                vec![6.0, -2.0, 1.0]];
    match gas(&mut matrix, &mut vector) {
        None => {
            println!("There os more then one solution");
        },
        Some(result) => {
            println!("{} {} {}", result[0], result[1], result[2]);
        }
    }
    return Ok(());
}