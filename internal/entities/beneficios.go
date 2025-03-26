package entities

type Beneficios struct {
	ID                    uint   `json:"id"`
	Categoria             string `json:"categoria"`              // Categoria do parceiro
	NomeParceiro          string `json:"nome_parceiro"`          // Nome do parceiro
	Thumbnail             string `json:"thumbnail"`              // Caminho ou URL para o thumbnail
	DescricaoCover        string `json:"descricao_cover"`        // Descrição da capa
	Img1                  []byte `json:"img1"`                   // Caminho ou URL para imagem 1
	Img2                  []byte `json:"img2,omitempty"`         // Caminho ou URL para imagem 2 (opcional)
	Img3                  []byte `json:"img3,omitempty"`         // Caminho ou URL para imagem 3 (opcional)
	InformacoesAdicionais string `json:"informacoes_adicionais"` // Breve descrição adicional
	MoreInfo              string `json:"more_info"`              // Informações adicionais
	Discounts             string `json:"discounts"`              // Informações sobre descontos
	Latitude              string `json:"latitude"`
	Longitude             string `json:"longitude"`
	LinkParceiro          string `json:"link_parceiro"`
	LinkWhatsapp          string `json:"link_whatsapp"`
	Validade              string `json:"validade"`
}
