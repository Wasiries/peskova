fn f(x: f64) -> f64 {
    return x.powf(3.0) - x - 2.0;
}

fn df(x: f64) -> f64 {
    return 3.0 * x.powf(2.0) - 1.0;
}


fn main() {
    let epsilon = 1e-8;
    let max_iter: usize = 1000;

    let mut x_old: f64 = 0.0;

    for _ in 0..max_iter {
        let derivative = df(x_old); 
        if derivative.abs() < 1e-10 {
            break;
        }
        let x_new = x_old - f(x_old) / derivative;

        if (x_new - x_old).abs() < epsilon {
            break;
        }
        x_old = x_new;
    }

    println!("approximation to solution: {}", x_old);
}