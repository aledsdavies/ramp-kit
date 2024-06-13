package models

type MetaInfo struct {
    Title       string
    Description string
    Keywords    []string
    Author      string
}


// WithTitle sets the Title field and returns the updated MetaInfo pointer.
func (m *MetaInfo) WithTitle(title string) *MetaInfo {
    m.Title = title
    return m
}

// WithDescription sets the Description field and returns the updated MetaInfo pointer.
func (m *MetaInfo) WithDescription(description string) *MetaInfo {
    m.Description = description
    return m
}

// WithKeywords sets the Keywords field and returns the updated MetaInfo pointer.
func (m *MetaInfo) WithKeywords(keywords []string) *MetaInfo {
    m.Keywords = keywords
    return m
}

// WithAuthor sets the Author field and returns the updated MetaInfo pointer.
func (m *MetaInfo) WithAuthor(author string) *MetaInfo {
    m.Author = author
    return m
}

