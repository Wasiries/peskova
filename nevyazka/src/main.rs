mod algo;
use std::rc::Rc;

fn main() {
    let temp: Vec<Vec<f64>> = vec![
        vec![5.0, 1.0, 0.0, 0.0, 0.0],
        vec![1.0, 5.0, 1.0, 0.0, 0.0],
        vec![0.0, 1.0, 5.0, 1.0, 0.0],
        vec![0.0, 0.0, 1.0, 5.0, 1.0],
        vec![0.0, 0.0, 0.0, 1.0, 5.0]
    ];
    let vector: Vec<f64> = vec![7.0, 14.0, 21.0, 28.0, 29.0];
    let epsilon: f64 = 1e-15;
    let matrix: Rc<Vec<Vec<f64>>> = Rc::new(temp);
    let m: algo::SquareMatrix = algo::SquareMatrix::make_matrix(5, matrix).unwrap();
    let x = m.equals_to(&vector, epsilon).unwrap();
    println!("{:#?}", x);
}
