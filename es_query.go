package elasticsearch

type EsQuery interface {
    Source() interface{}
}
