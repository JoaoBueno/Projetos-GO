package tipos

type Empresa struct {
	NI                   string              `json:"ni"`
	Porte                string              `json:"porte"`
	Endereco             Endereco            `json:"endereco"`
	Telefones            []Telefone          `json:"telefones"`
	DataAbertura         string              `json:"dataAbertura"`
	NomeFantasia         string              `json:"nomeFantasia"`
	CapitalSocial        int                 `json:"capitalSocial"`
	CnaePrincipal        Cnae                `json:"cnaePrincipal"`
	CnaeSecundarias      []Cnae              `json:"cnaeSecundarias"`
	NomeEmpresarial      string              `json:"nomeEmpresarial"`
	NaturezaJuridica     NaturezaJuridica    `json:"naturezaJuridica"`
	SituacaoEspecial     string              `json:"situacaoEspecial"`
	CorreioEletronico    string              `json:"correioEletronico"`
	SituacaoCadastral    SituacaoCadastral   `json:"situacaoCadastral"`
	MunicipioJurisdicao  MunicipioJurisdicao `json:"municipioJurisdicao"`
	TipoEstabelecimento  string              `json:"tipoEstabelecimento"`
	DataSituacaoEspecial string              `json:"dataSituacaoEspecial"`
}

type Endereco struct {
	UF             string    `json:"uf"`
	CEP            string    `json:"cep"`
	Pais           Pais      `json:"pais"`
	Bairro         string    `json:"bairro"`
	Numero         string    `json:"numero"`
	Municipio      Municipio `json:"municipio"`
	Logradouro     string    `json:"logradouro"`
	Complemento    string    `json:"complemento"`
	TipoLogradouro string    `json:"tipoLogradouro"`
}

type Pais struct {
	Codigo    string `json:"codigo"`
	Descricao string `json:"descricao"`
}

type Municipio struct {
	Codigo    string `json:"codigo"`
	Descricao string `json:"descricao"`
}

type Telefone struct {
	DDD    string `json:"ddd"`
	Numero string `json:"numero"`
}

type Cnae struct {
	Id        string `json:"codigo"`
	Descricao string `json:"descricao"`
}

type NaturezaJuridica struct {
	Codigo    string `json:"codigo"`
	Descricao string `json:"descricao"`
}

type SituacaoCadastral struct {
	Data   string `json:"data"`
	Codigo string `json:"codigo"`
	Motivo string `json:"motivo"`
}

type MunicipioJurisdicao struct {
	Codigo    string `json:"codigo"`
	Descricao string `json:"descricao"`
}
