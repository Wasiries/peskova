use std::{rc::Rc, fmt};

#[derive(Debug)]
pub enum MatrixError {
    VectorNotMatch,
    IncorrectMatrixStructure,
    UnsymmetricalMatrix,
    UnsolvableMatrix,
}

impl fmt::Display for MatrixError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            MatrixError::VectorNotMatch => {
                return write!(f, "Vector size not match matrix size");
            },
            MatrixError::IncorrectMatrixStructure => {
                return write!(f, "Incorrect matrix structure");
            },
            MatrixError::UnsymmetricalMatrix => {
                return write!(f, "Unsymmetrical matrix");
            },
            MatrixError::UnsolvableMatrix => {
                return write!(f, "Unsolvable matrix");
            },
        }
    }
}

impl std::error::Error for MatrixError {
    fn cause(&self) -> Option<&dyn std::error::Error> {
        return None;
    }
    fn description(&self) -> &str {
        match self {
            MatrixError::VectorNotMatch => {
                return "Vector size not match matrix size";
            },
            MatrixError::IncorrectMatrixStructure => {
                return "Incorrect matrix structure";
            },
            MatrixError::UnsymmetricalMatrix => {
                return "Unsymmetrical matrix";
            },
            MatrixError::UnsolvableMatrix => {
                return "Unsolvable matrix";
            },
        }
    }
    fn source(&self) -> Option<&(dyn std::error::Error + 'static)> {
        return None;
    }
}

pub struct SquareMatrix {
    size: usize,
    matrix: Rc<Vec<Vec<f64>>>
}

impl SquareMatrix {
    pub fn make_matrix(size_: usize, matrix_: Rc<Vec<Vec<f64>>>) -> Result<Self, Box<dyn std::error::Error>> {
        if size_ != matrix_.len() {
            return Err(Box::new(MatrixError::IncorrectMatrixStructure));
        }
        for item in matrix_.iter() {
            if size_ != item.len() {
                return Err(Box::new(MatrixError::IncorrectMatrixStructure));
            }
        }
        let this: SquareMatrix = SquareMatrix { size: size_, matrix: Rc::clone(&matrix_) };
        return Ok(this);
    }
    pub fn multiply(&self, vector: &Vec<f64>) -> Result<Vec<f64>, Box<dyn std::error::Error>> {
        if vector.len() != self.size {
            return Err(Box::new(MatrixError::VectorNotMatch));
        }
        let answer: Vec<f64> = self.matrix
        .iter()
        .map(|row| {
            row
            .iter()
            .enumerate()
            .map(|(position, value)| {
                *value * vector[position]
            })
            .sum()
        })
        .collect();

        return Ok(answer);
    }
    fn is_symmetrical(&self) -> bool {
        // for i in 0..self.matrix.len() {
        //     for j in 0..i {
        //         if self.matrix[i][j] != self.matrix[j][i] {
        //             return false;
        //         }
        //     }
        // }
        let value = self.matrix
        .iter()
        .enumerate()
        .all(|(position, row)| {
            row
            .iter()
            .enumerate()
            .filter(|(index, _)| {
                *index < position
            })
            .all(|(index, v)| {
                *v == self.matrix[index][position]
            })
        });

        return value;
    }
    fn norm(vector: &Vec<f64>) -> f64 {
        let answer: f64 = vector
        .iter()
        .map(|value|{ 
            value * value
        })
        .sum();

        return answer.powf(0.5);
    }
    pub fn equals_to(&self, vector: &Vec<f64>, epsilon: f64) -> Result<Vec<f64>, Box<dyn std::error::Error>> {
        let mut x: Vec<f64> = Vec::new();
        if vector.len() != self.size {
            return Err(Box::new(MatrixError::VectorNotMatch));
        }
        if !self.is_symmetrical() {
            return Err(Box::new(MatrixError::UnsymmetricalMatrix));
        }
        for _ in 0..self.size {
            x.push(0.0);
        }
        let mut max_iterations: usize = 1000000;
        let mut ax: Vec<f64> = self.multiply(&x)?;
        let mut y: Vec<f64> = ax
        .iter()
        .enumerate()
        .map(|(index, value)| {
            value - vector[index]
        })
        .collect();

        while SquareMatrix::norm(&y) > epsilon {
            let ay = self.multiply(&y)?;

            let numerator: f64 = y
            .iter()
            .enumerate()
            .map(|(index, element)| {
                element * ay[index]
            })
            .sum();

            let denominator: f64 = ay
            .iter()
            .map(|element|{ 
                element * element 
            })
            .sum();

            if denominator.abs() < 1e-15 {
                break;
            }

            let tau: f64 = numerator / denominator;
            x = x
            .iter()
            .enumerate()
            .map(|(index, element)| {
                element - tau * y[index]
            })
            .collect();

            ax = self.multiply(&x)?;
            y = ax
            .iter()
            .enumerate()
            .map(|(index, element)| {
                element - vector[index]
            })
            .collect();

            max_iterations -= 1;
            if max_iterations == 0 {
                return Err(Box::new(MatrixError::UnsolvableMatrix));
            }
        }
        return Ok(x);
    }
}