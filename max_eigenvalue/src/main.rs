mod algorithm;

fn main() {
    let matrix: Vec<Vec<f64>> = vec![
        vec![1.0, 2.0, 3.0, 4.0, 5.0, 6.0],
        vec![1.0, 2.0, 3.0, 4.0, 5.0, 6.0],
        vec![1.0, 2.0, 3.0, 4.0, 5.0, 6.0],
        vec![1.0, 2.0, 3.0, 4.0, 5.0, 6.0],
        vec![1.0, 2.0, 3.0, 4.0, 5.0, 6.0],
        vec![1.0, 2.0, 3.0, 4.0, 5.0, 6.0],
    ];
    let answer = algorithm::power_iteration(&matrix, 1000, 1e-10).unwrap();
    println!("max eigenvalue: {}", answer);
}