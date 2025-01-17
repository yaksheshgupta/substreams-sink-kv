mod pb;
use std::{str, thread, time};
#[allow(unused_imports)]
use wasmedge_bindgen::*;
use wasmedge_bindgen_macro::*;
use crate::pb::test::{
    TestGetRequest,
    TestPrefixRequest,
    TestScanRequest,
    TestGetManyRequest,
    TestSleepRequest,
    TestPanicRequest,
    Tuple,
    Tuples,
    Response
};
use prost::Message;
use substreams_sink_kv::pb::sf::substreams::sink::types::v1::KvPair;
use substreams_sink_kv::prelude::*;

#[wasmedge_bindgen]
pub fn sf_test_v1_testservice_testget(v: Vec<u8>) -> Result<Vec<u8>, String> {
    let req = TestGetRequest::decode(&v[..]).expect("Failed to decode");
    let store = Store::new();

    let opt_value = store.get(&req.key);

    match opt_value {
        Some(kv_pair)=> {
            Ok(to_key_value(&kv_pair).encode_to_vec())
        },
        None => {
            Err("not found".to_string())
        }
    }
}

#[wasmedge_bindgen]
pub fn sf_test_v1_testservice_testgetmany(v: Vec<u8>) -> Result<Vec<u8>, String> {
    let req = TestGetManyRequest::decode(&v[..]).expect("Failed to decode");
    let store = Store::new();

    let mut keys: Vec<String> = vec![];
    for k in req.keys {
        keys.push(k.to_string())
    }

    let opt_kv_pairs = store.get_many(keys);

    match opt_kv_pairs {
        Some(kv_pairs)=> {
            let mut response = Tuples{ pairs: vec![]};
            for kv_pair in kv_pairs.pairs {
                response.pairs.push(to_key_value(&kv_pair));
            }
            Ok(response.encode_to_vec())
        },
        None => {
            Err("not found".to_string())
        }
    }
}

#[wasmedge_bindgen]
pub fn sf_test_v1_testservice_testprefix(v: Vec<u8>) -> Result<Vec<u8>, String> {
    let req = TestPrefixRequest::decode(&v[..]).expect("Failed to decode");
    let store = Store::new();

    let mut limit : Option<u32> = None;
    if let Some(l) = req.limit {
        limit = Some(l as u32)
    }

    let kv_pairs = store.prefix(&req.prefix, limit);

    let mut response = Tuples{ pairs: vec![]};
    for kv_pair in kv_pairs.pairs {
        response.pairs.push(to_key_value(&kv_pair));
    }
    return Ok(response.encode_to_vec());
}

#[wasmedge_bindgen]
pub fn sf_test_v1_testservice_testscan(v: Vec<u8>) -> Result<Vec<u8>, String> {
    let req = TestScanRequest::decode(&v[..]).expect("Failed to decode");
    let store = Store::new();

    let mut limit : Option<u32> = None;
    if let Some(l) = req.limit {
        limit = Some(l as u32)
    }

    let kv_pairs = store.scan(&req.start,&req.exclusive_end, limit);

    let mut response = Tuples{ pairs: vec![]};
    for kv_pair in kv_pairs.pairs {
        response.pairs.push(to_key_value(&kv_pair));
    }
    return Ok(response.encode_to_vec());
}

#[wasmedge_bindgen]
pub fn sf_test_v1_testservice_testsleep(v: Vec<u8>) -> Result<Vec<u8>, String> {
    let req = TestSleepRequest::decode(&v[..]).expect("Failed to decode");
    println!("START {}", req.request_id);
    let ten_millis = time::Duration::from_millis(req.duration as u64);

    thread::sleep(ten_millis);

    let resp = Response{
        output: format!("{} completed",req.request_id),
    };
    println!("END {}", req.request_id);
    Ok(resp.encode_to_vec())
}


#[wasmedge_bindgen]
pub fn sf_test_v1_testservice_testpanic(v: Vec<u8>) -> Result<Vec<u8>, String> {
    substreams_sink_kv::error::register_panic_hook();

    let req = TestPanicRequest::decode(&v[..]).expect("Failed to decode");

    if req.should_panic {
        panic!{"panicking"}
    }

    println!("you should not see this");
    let resp = Response{
        output: "did not panic".to_string(),
    };
    Ok(resp.encode_to_vec())
}


pub fn to_key_value(kv_pair: &KvPair) -> Tuple {
    let output = str::from_utf8(&*kv_pair.value).unwrap();
    Tuple{
        key: kv_pair.key.clone(),
        value: output.to_string(),
    }
}
