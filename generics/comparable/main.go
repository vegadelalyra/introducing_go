func Contains[T comparable](slice []T, target T) bool {
    for _, item := range slice {
        if item == target { // This '==' only works because of 'comparable'
            return true
        }
    }
    return false
}
