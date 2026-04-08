pub fn power_iteration(matrix: &Vec<Vec<f64>>, max_iter: usize, epsilon: f64) -> Option<f64> {
    let mut b = vec![1.0; matrix.len()];
    let mut lambda_old = 0.0;

    for _ in 0..max_iter {
        let mut b_new: Vec<f64> = matrix
        .iter()
        .enumerate()
        .map(|(_, value)| {
            value
            .iter()
            .zip(&b)
            .map(|(first, second)| {
                first * second
            })
            .sum()
        })
        .collect();


        let b_bnew: f64 = b
        .iter()
        .zip(&b_new)
        .map(|(first, second)| {
            first * second
        })
        .sum();
        let b_b: f64 = b
        .iter()
        .map(|value| {
            value * value 
        })
        .sum();

        if b_b.abs() < 1e-10 {
            return None;
        }
        let lambda = b_bnew / b_b;

        let mut norm: f64 = b_new
        .iter()
        .map(|value| {
            value * value
        })
        .sum();
        norm = norm.sqrt();

        if norm < 1e-12 {
            return None;
        }
        
        b_new = b_new
        .iter()
        .map(|value| {
            value / norm
        })
        .collect();

        b = b_new;

        if (lambda - lambda_old).abs() < epsilon {
            return Some(lambda);
        }
        lambda_old = lambda;
    }
    return None;
}