
use std::{rc::Rc};

#[derive(Debug)]
pub struct ZeroDivisionError;

impl std::fmt::Display for ZeroDivisionError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        return write!(f, "Division by zero");
    }
}

impl std::error::Error for ZeroDivisionError {
    fn cause(&self) -> Option<&dyn std::error::Error> {
        return None;
    }
    fn description(&self) -> &str {
        return "Division by zero";
    }
    fn source(&self) -> Option<&(dyn std::error::Error + 'static)> {
        return None;
    }
}

pub fn norm(vector: &Vec<f64>) -> f64 {
    let answer: f64 = vector
    .iter()
    .map(|value| {
        value * value
    })
    .sum();
    return answer.powf(0.5);
}


pub fn conjugate_gradient(matrix: Rc<Vec<Vec<f64>>>, vector: Vec<f64>, epsilon: f64) -> Result<Vec<f64>, Box<dyn std::error::Error>> {
    let size = vector.len();
    let mut x: Vec<f64> = Vec::new();
    for _ in 0..size {
        x.push(0.0);
    }

    let mut r: Vec<f64> = vector
    .iter()
    .enumerate()
    .map(|(index, element)| {
        let t: f64 = x
        .iter()
        .enumerate()
        .map(|(position, value)| {
            value * matrix[index][position]
        })
        .sum();
        element - t
    })
    .collect();

    let mut p = r.clone();

    while norm(&r) > epsilon {
        let a_numerator: f64 = r
        .iter()
        .map(|element| {
            element * element 
        })
        .sum();

        let ap: Vec<f64> = p
        .iter()
        .enumerate()
        .map(|(index, _)| {
            p
            .iter()
            .enumerate()
            .map(|(position, value)| {
                value * matrix[index][position]
            })
            .sum() 
        })
        .collect();

        let a_denominator: f64 = p
        .iter()
        .zip(&ap)
        .map(|(first, second)| {
            first * second
        })
        .sum();

        if a_denominator.abs() < epsilon {
            return Err(Box::new(ZeroDivisionError));
        }

        let alpha = a_numerator / a_denominator;
        x = x
        .iter()
        .zip(&p)
        .map(|(first, second)| {
            first + alpha * second
        })
        .collect();

        let r_next: Vec<f64> = r
        .iter()
        .zip(&ap)
        .map(|(first, second)| {
            first - alpha * second
        })
        .collect();
        
        let b_numerator: f64 = r_next
        .iter()
        .map(|value| {
            value * value
        })
        .sum();
        let b_denominator: f64 = r
        .iter()
        .map(|value| {
            value * value
        })
        .sum();

        if b_denominator.abs() < epsilon {
            return Err(Box::new(ZeroDivisionError));
        }

        let beta = b_numerator / b_denominator;
        r = r_next;
        p = r
        .iter()
        .zip(&p)
        .map(|(first, second)| {
            first + beta * second
        })
        .collect();
    }
    return Ok(x);
}