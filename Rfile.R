install.packages("microbenchmark")
library(microbenchmark)
library(keras)

# Load MNIST dataset
data <- dataset_mnist()
train_images <- data$train$x
train_labels <- data$train$y
test_images <- data$test$x
test_labels <- data$test$y

# Preprocess the data
train_images <- array_reshape(train_images, c(nrow(train_images), 28 * 28)) / 255
test_images <- array_reshape(test_images, c(nrow(test_images), 28 * 28)) / 255

# Define and compile the model
model <- keras_model_sequential() %>%
  layer_dense(units = 512, activation = 'relu', input_shape = c(28 * 28)) %>%
  layer_dense(units = 10, activation = 'softmax') %>%
  compile(
    optimizer = 'rmsprop',
    loss = 'sparse_categorical_crossentropy',
    metrics = c('accuracy')
  )

# Benchmark model training
benchmark_result <- microbenchmark(
  model %>% fit(train_images, train_labels, epochs = 5, batch_size = 128),
  times = 10  # Number of times to run the benchmark
)

# Print benchmark results
print(benchmark_result)
