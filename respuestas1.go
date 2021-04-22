import(
	"encoding/json"
	"fmt"
)
//Estructuras base
//Estructura de tarjeta con id
type Card struct{
	Id string
	Type string
	Brand string
	Address string
	Number string
	Holder_name string
	Expiration_year string
	Expiration_month string
	Allows_charges bool
	Allows_payouts bool
	Creation_date string
	Bank_name string
	Bank_code string
}

//Estructura de tarjeta para confirmar cargo
type Card_charge struct{
	Type string
	Brand string
	Address string				json:"address,omitempty"
	Card_number string
	Holder_name string
	Expiration_year string		json:"expiration_year,omitempty"
	Expiration_month string		json:"expiration_month,omitempty"
	Allows_charges bool
	Allows_payouts bool
	Bank_name string
	Bank_code string
}

//Estructura de tarjeta para suscripciones//PAQUE???? la respuesta de sus usa la estructura de arriba
type Card_subs struct{
	Type string
	Brand string
	Card_number string
	Holder_name string
	Expiration_year string
	Expiration_month string
	Allows_charges bool
	Allows_payouts bool
	Creation_date string
	Bank_name string
	Customer_id string	json:"customer_id,omitempty"
	Bank_code string
}

//Estructura de tipo de cambio
type Exchange_rate struct{
	From string
	Date string
	Value float64
	To string
}

//Estructura metodo de pago con redirección
type Payment_method_redi struct{
	Type string
	Url string
}

//Estructura customer
type Customer struct{
	Name string
	Last_name string
	Email string
	Phone_number string
	Creation_date string
	Clabe string
	External_id string
}

//Estructura metodo de pago en tienda
type Payment_method_tienda struct{
	Type string
	Reference string
	Barcode_url string
	Barcode_paybin_url string
}

//Estructura metodo de pago banco
type Payment_method_bank struct{
	Type string
	Agreement string
	Bank string
	Clabe string
	Name string
}

//Estructura fee
type Fee struct{
	Amount float64
	Tax int
	Currency string
}


//Estructura devolucion
type Refund_charge struct{
	Id string 
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Operation_type string
	Transaction_type string
	Status string
	Currency string
	Creation_date string
	Operation_date string 
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string 
	Customer_id string 
}

//Estructura de cuenta de banco con ID(no se si con solo omitir o si se requiere otra estructura en estos casos)
type Bank_account_id struct{
	Id string		json:"id,omitempty"
	Clable string
	Bank_code string
	Bank_name string
	Alias string	json:"alias,omitempty"
	Holder_name string
}
//Estructura de cuenta de banco
type Bank_account_ struct{
	Clable string
	Bank_code string
	Bank_name string
	Alias string	json:"alias,omitempty"
	Holder_name string
}

//Estructura cuenta de banco RFC
type Bank_account_rfc struct{
	Clabe string
	Bank_code string
	Bank_name string
	Rfc string
	Holder_name string
}

//dirección de objeto cliente 
type Address_client struct{
	Line1 string
	Line2 string
	Line3 string
	Postal_code string
	State string
	City string
	Country_code string
}

// store de objeto cliente
type Store_client struct{
	Reference string
	Barcode_url string
}
/////////////////////////////////////////////////////////////
//CARGOS
//Respuesta al cargo con id o token
type ResCargoIdTok struct{
	Id string
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Operation_type string
	Transaction_type string
	Card Card
	Status string
	Currency string
	Exchange_rate Exchange_rate
	Creation_date string
	Operation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string
}

//Respuesta al cargo con redireccionamiento
type ResCargoRedirect struct{
	Id string
	Authorization string	json:"authorization,omitempty"
	Method string
	Operation_type string
	Transaction_type string
	Status string
	Conciliated bool
	Creation_date string
	Operation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Amount float64
	Currency string
	Payment_method Payment_method_redi
	Customer Customer
}

//Respuesta a cargo en tienda
type ResCargoTienda struct{
	Id string
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Operation_type string
	Transaction_type string
	Status string
	Currency string
	Creation_date string
	Operation_date string
	Due_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string
	Customer_id string
	Payment_method Payment_method_tienda
}

//Respuesta a cargo en banco
type ResCargoBanco struct{
	Id string
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Operation_type string
	Transaction_type string
	Status string
	Currency string
	Creation_date string
	Operation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string
	Customer_id string
	Payment_method Payment_method_bank
}

//Respuesta a cargo en Alipay
type ResCargoAlipay struct{
	Id string
	Authorization string	json:"authorization,omitempty"
	Operation_type string
	Method string
	Transaction_type string
	Status string
	Conciliated bool
	Creation_date string
	Operation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string
	Due_date string
	Payment_method Payment_method_redi
	Amount float64
	Currency string
	Fee Fee
}

//Estructura de confirmación de cargo
type ResConfirmaCargo struct{
	Id string
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Operation_type string
	Transaction_type string
	Card Card_charge
	Status string
	Currency string
	Creation_date string
	Operation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string
	Customer_id string
}

//Estructura de devolución u obtención de cargo (misma estructura, distinto fin)
type ResDevuelveCargo struct{
	Id string
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Operation_type string
	Transaction_type string
	Card Card_charge
	Status string
	Refund Refund_charge
	Currency string
	Creation_date string
	Operation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string
	Customer_id string
}

//Estructuras lista de cargos(Ambos van en un arreglo)
type ResListaCargos1 struct{
	Id string
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Operation_type string
	Transaction_type string
	Card Card_charge
	Status string
	Currency string
	Creation_date string
	Operation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string 
	Customer_id string
}
type ResListaCargos2 struct{
	Id string
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Operation_type string
	Transaction_type string
	Status string
	Currency string
	Creation_date string
	Operation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string 
	Customer_id string
	Payment_method Payment_method_bank
}


//Pagos o retiros

//Pago a id registrado
type ResPagoId struct{
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Currency string
	Operation_type string
	Transaction_type string
	Bank_account Bank_account_id
	Status string
	Id string
	Creation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string
	Customer_id string
}

//Pago por cuenta bancaria
type ResPagoCB struct{
	Id string
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Operation_type string
	Transaction_type string
	Status string
	Currency string
	Creation_date string 
	Operation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string
	Bank_account Bank_account_
	Customer_id string
}

//Respuesta obtener un pago
type ResObtenerPago struct{
	Id string
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Operation_type string
	Transaction_type string
	Bank_account Bank_account_
	Status string
	Currency string
	Creation_date string 
	Operation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string
	Customer_id string
}

//Respuesta de Cancelar un pago // identico al anterior, se puede usar el mismo?
type ResCancelPago struct{
	Id string
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Operation_type string
	Transaction_type string
	Bank_account Bank_account_
	Status string
	Currency string
	Creation_date string 
	Operation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string
	Customer_id string
}


//Listado de pagos (en array)
//Respuesta 1 a listas de pagos
type ResListadoPagos struct{
	Id string
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Operation_type string
	Transaction_type string
	Status string
	Currency string
	Creation_date string
	Operation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string 
	Bank_account Bank_account_
	Customer_id string
}

//Respuesta resumen de pagos
type ResResumenPagos struct{
	In int64
	Out int64
	Charged_adjustaments int64
	Refunded_adjustaments int64
}

//Respuestas detalle de pagos (en array)
//Respuesta1 detalles de pagos 
type ResDetPagosBank struct{
	Id string
	Authorization string
	Method string
	Operation_type string
	Transaction_type string
	Status string
	Conciliated string	json:"conciliated,omitempty"
	Creation_date string
	Operation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string	json:"order_id,omitempty"
	Bank_account Bank_account_rfc
	Amount float64
	Currency string
}
//Respuesta2 detalles de pagos
type ResDetPagosCard struct{
	Id string
	Authorization string
	Method string
	Operation_type string
	Transaction_type string
	Card Card_charge
	Status string
	Conciliated string	json:"conciliated,omitempty"
	Creation_date string
	Operation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string	json:"order_id,omitempty"
	Amount float64
	Currency string
}

///////////////////////////////////////////////////////////////
//Clientes 
//Objeto cliente
type Cliente struct{
	Id string
	Creation_date string
	Name string
	Last_name string
	Email string
	Phone_number string
	External_id string
	Status string
	Balance float64
	Address Address_client
	Store Store_client
	Clabe string
}

//Respuesta Nuevo cliente
type ResNuevoCliente struct{
	Id string
	Name string
	Last_name string	json:"last_name,omitempty"
	Email string
	Phone_number string	json:"phone_number,omitempty"
	Address string	json:"address,omitempty"
	Creation_date string
	External_id string	json:"external_id,omitempty"
	Store Store_client
	Clabe string
}

//Respuesta actualiza cliente //En teoria igual a la anterior, ¿la respuesta nulla es comparable al tipo estructura?// Obtener cliente existente exactamente la misma estructura, lo dejo para ambas
type ResActualizaObtieneCliente struct{
	Id string
	Name string
	Last_name string	json:"last_name,omitempty"
	Email string
	Phone_number string	json:"phone_number,omitempty"
	Address Address_client
	Store Store_client
	Clabe string
	Creation_date string
	External_id string	json:"external_id,omitempty"
}

//Respuesta listado de clientes//Arreglo de objetos
type ResListaCliente struct{
	Id string
	Creation_date string
	Name string
	Last_name string
	Email string
	Phone_number string
	Status string
	Store Store_client
	Clabe string 
}

//////////////////////////////////////////////////////////////
//Transferencias (por si si o por si no)
//Transferencia entre clientes 
type ResTransfClientes struct{
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Operation_type string
	Currency string
	Transaction_type string
	Status string
	Id string
	Creation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string
	Customer_id string
	Store Store_client
	Clabe string
}

//Obtener una transferencia
type ResObtieneTransf struct{
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Operation_type string
	Currency string
	Transaction_type string
	Status string
	Id string
	Creation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string
	Customer_id string
}

//Listado de transferencias//Listado de objetos//Identica a la anterior
type ResListaTransf struct{
	Amount float64
	Authorization string	json:"authorization,omitempty"
	Method string
	Currency string
	Operation_type string
	Transaction_type string
	Status string
	Id string
	Creation_date string
	Description string
	Error_message string	json:"error_message,omitempty"
	Order_id string
	Customer_id string
}

//////////////////////////////////////////////////////////////
//Tarjetas???
//Objeto tarjeta
type ObjTarjeta struct{
	Type string
	Brand string
	Address Address_client
	Id string
	Card_number string
	Holder_name string
	Expiration_year string
	Expiration_month string
	Allows_charges bool
	Allows_payouts bool
	Creation_date string
	Bank_name string
	Bank_code string
	Customer_id string
	Points_card bool
}

//Respuesta Crear tarjeta
type ResCreaTarjeta struct{
	Id string
	Type string
	Brand string
	Card_number string
	Holder_name string
	Expiration_year string
	Expiration_month string
	Allows_charges bool
	Allows_payouts bool
	Creation_date string
	Bank_name string
	Bank_code string
	Customer_id string
	Points_card bool
}

//Respuesta crear tarjeta token
type ResCreaTarjetaToken struct{
	Type string
	Brand string
	Id string
	Card_number string
	Holder_name string
	Expiration_year string
	Expiration_month string
	Allows_charges bool
	Allows_payouts bool
	Creation_date string
	Bank_name string
	Bank_code string
	Customer_id string
}

//Respuesta obtener tarjeta //Identico a crear tarjeta con token, lo dejo por diferenciación?
type ResObtieneTarjeta struct{
	Id string
	Type string
	Brand string
	Card_number string
	Holder_name string
	Expiration_year string
	Expiration_month string
	Allows_charges bool
	Allows_payouts bool
	Creation_date string
	Bank_name string
	Bank_code string
	Customer_id string
}

//Respuesta consulta de puntos //En arreglo
type ResConsPuntos struct{
	Points_type string
	Remaining_points string
	Remaining_mxn string//?? la documentación lo maneja como cadena, sería bueno cambiarlo a float64?
}

//Respuesta listado de tarjetas//Arreglo de objetos
type ResListaTarj struct{
	Id string
	Card_number string
	Holder_name string
	Expiration_year string
	Expiration_month string
	Allows_charges bool
	Allows_payouts bool
	Creation_date string
	Bank_name string
	Bank_code string
	Type string
	Brand string
}
//Actualizar tarjeta regresa un objeto vacío

//////////////////////////////////////////////////////////////
//Cuentas bancarias
//Objeto cuenta bancaria
type ObjCB struct{
	Id string
	Clabe string
	Bank_code string
	Bank_name string
	Alias string	json:"alias,omitempty"
	Holder_name string
	Creation_date string
}

//Respuesta crear CB //Identico al anterior//Aplica para obtener
type ResCreaObtieneCB struct{
	Id string
	Clabe string
	Bank_code string
	Bank_name string
	Alias string
	Holder_name string
	Creation_date string
}

//La lista es un arreglo de ObjCB
//////////////////////////////////////////////////////////////
//PLANES
//Objeto plan
type ObjPlan struct{
	Name string
	Status string
	Amount float64
	Currency string
	Id string
	Creation_date string
	Repeat_every float64
	Repeat_unit string
	Retry_times float64
	Status_after_retry string
	Trial_days int
}

//Respuesta Nuevo plan //Identico al OBJ, actualizar plan,Obtener plan y el listado es un arreglo de estos objetos YOLO
type ResNuevoPlan struct{
	Id string
	Name string
	Status string
	Amount float64
	Currency string
	Creation_date string
	Repeat_every float64
	Repeat_unit string
	Retry_times float64
	Status_after_retry string
	Trial_days int
}

//////////////////////////////////////////////////////////////
//Suscripciones
//Objeto suscripción
type ObjSuscripcion struct{
	Status string
	Card Card_subs
	Id string
	Cancel_at_period_end bool
	Charge_date string
	Creation_date string
	Current_period_number int
	Period_end_date string
	Trial_end_date string
	Plan_id string
	Customer_id string
}

//Respuesta crea nueva suscripcion //Identica al anterior salvo por el tipo de tarjeta
type ResNuevaSus struct{ 
	Id string
	Status string
	Card Card_charge
	Cancel_at_period_end bool
	Charge_date string
	Creation_date string
	Current_period_number int
	Period_end_date string
	Trial_end_date string
	Plan_id string
	Customer_id string
}