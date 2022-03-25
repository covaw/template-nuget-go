// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Pedido.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Pedido struct {
	Id string `json:"id"`

	NumeroDePedido int32 `json:"numeroDePedido"`

	CicloDelPedido string `json:"cicloDelPedido"`

	CodigoDeContratoInterno int64 `json:"codigoDeContratoInterno"`

	EstadoDelPedido string `json:"estadoDelPedido"`

	CuentaCorriente int64 `json:"cuentaCorriente"`

	Cuandos string `json:"cuandos"`
}

const PedidoAvroCRC64Fingerprint = "\xe6\xaa\bO<\xe4bH"

func NewPedido() Pedido {
	r := Pedido{}
	return r
}

func DeserializePedido(r io.Reader) (Pedido, error) {
	t := NewPedido()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePedidoFromSchema(r io.Reader, schema string) (Pedido, error) {
	t := NewPedido()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePedido(r Pedido, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.NumeroDePedido, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CicloDelPedido, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.CodigoDeContratoInterno, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EstadoDelPedido, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.CuentaCorriente, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Cuandos, w)
	if err != nil {
		return err
	}
	return err
}

func (r Pedido) Serialize(w io.Writer) error {
	return writePedido(r, w)
}

func (r Pedido) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"numeroDePedido\",\"type\":\"int\"},{\"name\":\"cicloDelPedido\",\"type\":\"string\"},{\"name\":\"codigoDeContratoInterno\",\"type\":\"long\"},{\"name\":\"estadoDelPedido\",\"type\":\"string\"},{\"name\":\"cuentaCorriente\",\"type\":\"long\"},{\"name\":\"cuandos\",\"type\":\"string\"}],\"name\":\"Andreani.Scheme.Onboarding.Pedido\",\"type\":\"record\"}"
}

func (r Pedido) SchemaName() string {
	return "Andreani.Scheme.Onboarding.Pedido"
}

func (_ Pedido) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Pedido) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Pedido) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Pedido) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Pedido) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Pedido) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Pedido) SetString(v string)   { panic("Unsupported operation") }
func (_ Pedido) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Pedido) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Id}

		return w

	case 1:
		w := types.Int{Target: &r.NumeroDePedido}

		return w

	case 2:
		w := types.String{Target: &r.CicloDelPedido}

		return w

	case 3:
		w := types.Long{Target: &r.CodigoDeContratoInterno}

		return w

	case 4:
		w := types.String{Target: &r.EstadoDelPedido}

		return w

	case 5:
		w := types.Long{Target: &r.CuentaCorriente}

		return w

	case 6:
		w := types.String{Target: &r.Cuandos}

		return w

	}
	panic("Unknown field index")
}

func (r *Pedido) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Pedido) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Pedido) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Pedido) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Pedido) HintSize(int)                     { panic("Unsupported operation") }
func (_ Pedido) Finalize()                        {}

func (_ Pedido) AvroCRC64Fingerprint() []byte {
	return []byte(PedidoAvroCRC64Fingerprint)
}

func (r Pedido) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["numeroDePedido"], err = json.Marshal(r.NumeroDePedido)
	if err != nil {
		return nil, err
	}
	output["cicloDelPedido"], err = json.Marshal(r.CicloDelPedido)
	if err != nil {
		return nil, err
	}
	output["codigoDeContratoInterno"], err = json.Marshal(r.CodigoDeContratoInterno)
	if err != nil {
		return nil, err
	}
	output["estadoDelPedido"], err = json.Marshal(r.EstadoDelPedido)
	if err != nil {
		return nil, err
	}
	output["cuentaCorriente"], err = json.Marshal(r.CuentaCorriente)
	if err != nil {
		return nil, err
	}
	output["cuandos"], err = json.Marshal(r.Cuandos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Pedido) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroDePedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDePedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroDePedido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cicloDelPedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CicloDelPedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cicloDelPedido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoDeContratoInterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeContratoInterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigoDeContratoInterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["estadoDelPedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoDelPedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for estadoDelPedido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cuentaCorriente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CuentaCorriente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cuentaCorriente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cuandos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cuandos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cuandos")
	}
	return nil
}
